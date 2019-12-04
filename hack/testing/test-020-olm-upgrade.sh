#!/bin/bash
# Given an OLM manifest, verify a green field deployment
# of cluster logging by asserting CLO creates the resources
# that begets the operands that make up logging.

set -e
if [ "${DEBUG:-}" = "true" ]; then
	set -x
fi

repo_dir="$( cd "$(dirname "$0")/../.." ; pwd -P )"
source "$repo_dir/hack/testing/utils"
source "$repo_dir/hack/testing/assertions"

start_seconds=$(date +%s)

ARTIFACT_DIR=${ARTIFACT_DIR:-"$repo_dir/_output"}
if [ ! -d $ARTIFACT_DIR ] ; then
  mkdir -p $ARTIFACT_DIR
fi

KUBECONFIG=${KUBECONFIG:-$HOME/.kube/config}
TIMEOUT_MIN=$((2 * $minute))
NAMESPACE="openshift-logging"
manifest_dir=${repo_dir}/manifests
version=$(basename $(find $manifest_dir -type d | sort -r | head -n 1))
previous_version=$(echo $version | awk '{print $1 - 0.1}')

cleanup(){
  local return_code="$?"
  set +e
  log::info "Running cleanup"
  end_seconds=$(date +%s)
  runtime="$(($end_seconds - $start_seconds))s"
  oc -n openshift-operators-redhat -o yaml get subscription elasticsearch-operator > $ARTIFACT_DIR/subscription-eo.yml 2>&1 ||:
  oc logs -n ${NAMESPACE} deployment/cluster-logging-operator > $ARTIFACT_DIR/cluster-logging-operator.log 2>&1 ||:
  oc -n ${NAMESPACE} -o yaml get subscription cluster-logging-operator > $ARTIFACT_DIR/subscription-clo.yml 2>&1 ||:
  oc -n ${NAMESPACE} -o yaml get configmap cluster-logging > $ARTIFACT_DIR/configmap-for-catalogsource-clo.yml 2>&1 ||:
  oc -n ${NAMESPACE} -o yaml get catalogsource cluster-logging > $ARTIFACT_DIR/catalogsource-clo.yml 2>&1 ||:
  oc  -n openshift-operator-lifecycle-manager logs --since=$runtime deployment/catalog-operator > $ARTIFACT_DIR/catalog-operator.logs 2>&1 ||:
  oc  -n openshift-operator-lifecycle-manager logs --since=$runtime deployment/olm-operator > $ARTIFACT_DIR/olm-operator.logs 2>&1 ||:
  oc describe -n ${NAMESPACE} deployment/cluster-logging-operator > $ARTIFACT_DIR/cluster-logging-operator.describe.after_update  2>&1 ||:

  
  for item in "crd/elasticsearches.logging.openshift.io" "crd/clusterloggings.logging.openshift.io" "ns/openshift-logging" "ns/openshift-operators-redhat"; do
    oc delete $item --wait=true --ignore-not-found --force --grace-period=0
  done
  for item in "ns/openshift-logging" "ns/openshift-operators-redhat"; do
    try_until_text "oc get ${item} --ignore-not-found" "" "$((1 * $minute))"
  done

  exit ${return_code}
}
trap cleanup exit

if [ -n "${IMAGE_CLUSTER_LOGGING_OPERATOR:-}" ] ; then
  source "$repo_dir/hack/common"
fi
if [ -n "${IMAGE_FORMAT:-}" ] ; then
  IMAGE_CLUSTER_LOGGING_OPERATOR=$(sed -e "s,\${component},cluster-logging-operator," <(echo $IMAGE_FORMAT))
else
  IMAGE_CLUSTER_LOGGING_OPERATOR=${IMAGE_CLUSTER_LOGGING_OPERATOR:-registry.svc.ci.openshift.org/origin/${OCP_VERSION}:cluster-logging-operator}
fi


# deploy EO
deploy_marketplace_operator "openshift-operators-redhat" "elasticsearch-operator"  "$previous_version" "elasticsearch-operator" "true"

# verify operator is ready
try_until_text "oc -n openshift-operators-redhat get deployment elasticsearch-operator -o jsonpath={.status.availableReplicas} --ignore-not-found" "1" ${TIMEOUT_MIN}

# deploy CLO
deploy_marketplace_operator "openshift-logging" "cluster-logging-operator" "$previous_version" "cluster-logging"

# verify operator is ready
try_until_text "oc -n openshift-logging get deployment cluster-logging-operator -o jsonpath={.status.updatedReplicas} --ignore-not-found" "1" ${TIMEOUT_MIN}

# deploy cluster logging
oc -n $NAMESPACE create -f ${repo_dir}/hack/cr.yaml

assert_resources_exist

oc describe -n ${NAMESPACE} deployment/cluster-logging-operator > $ARTIFACT_DIR/cluster-logging-operator.describe.before_update 2>&1

deploy_config_map_catalog_source $NAMESPACE ${repo_dir}/manifests "${IMAGE_CLUSTER_LOGGING_OPERATOR}"
# #patch catalog source
# #olm.skipRange: ">=4.2.0 <4.3.0"
# range=">=$previous_version.0 <$(echo $version | awk '{print $1 + 0.1}').0"
# oc -n $NAMESPACE get configmap cluster-logging -o yaml | sed -e "s~olm.skipRange:.*~olm.skipRange: ${range}~" | oc replace -n $NAMESPACE -f -

# patch subscription
payload="{\"op\":\"replace\",\"path\":\"/spec/source\",\"value\":\"cluster-logging\"}"
payload="$payload,{\"op\":\"replace\",\"path\":\"/spec/sourceNamespace\",\"value\":\"openshift-logging\"}"
payload="$payload,{\"op\":\"replace\",\"path\":\"/spec/channel\",\"value\":\"$version\"}"
oc -n $NAMESPACE patch subscription cluster-logging-operator --type json -p "[$payload]"

#verify deployment is rolled out
try_until_text "oc -n openshift-logging get deployment cluster-logging-operator -o jsonpath={.spec.template.spec.containers[0].image}" "${IMAGE_CLUSTER_LOGGING_OPERATOR}" ${TIMEOUT_MIN}

# verify operator is ready
try_until_text "oc -n openshift-logging get deployment cluster-logging-operator -o jsonpath={.status.updatedReplicas} --ignore-not-found" "1" ${TIMEOUT_MIN}

assert_resources_exist
