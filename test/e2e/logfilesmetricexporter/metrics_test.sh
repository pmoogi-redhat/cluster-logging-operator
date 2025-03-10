#!/usr/bin/bash

#ensure you are logged into OCP cluster beforehand
#ensure clo operator is installed and CL, CLF instances are running beforehand

export NAMESPACE=openshift-logging


declare -a metrics=("log_logged_bytes_total")

collectorpod=`oc get pods -n openshift-logging | grep collector | awk 'NR==1{print $1}' `
token=`oc sa get-token prometheus-k8s -n openshift-monitoring`

# ## now loop through the above array
for metricname in "${metrics[@]}"

do 
  echo "$metricname"
  count=`oc exec -n openshift-logging ${collectorpod} -c logfilesmetricexporter -- curl -k -H "Authorization: Bearer ${token}" -s -H "Content-type: application/json" https://collector.openshift-logging.svc:2112/metrics | grep -s -c ${metricname}`

  if [[ $count -ge 1 ]]
  then 
    echo "metric found $metricname"
  else 
    echo "metric not found $metricname"
    exit 1
  fi

done