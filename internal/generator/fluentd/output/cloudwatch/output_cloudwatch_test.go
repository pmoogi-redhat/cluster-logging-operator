package cloudwatch

import (
	"path"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"

	loggingv1 "github.com/openshift/cluster-logging-operator/apis/logging/v1"
	"github.com/openshift/cluster-logging-operator/internal/constants"
	"github.com/openshift/cluster-logging-operator/internal/generator"

	. "github.com/openshift/cluster-logging-operator/test/matchers"
)

var _ = Describe("Generating fluentd config", func() {
	var (
		g      generator.Generator
		output = loggingv1.OutputSpec{
			Type: loggingv1.OutputTypeCloudwatch,
			Name: "my-cloudwatch",
			OutputTypeSpec: loggingv1.OutputTypeSpec{
				Cloudwatch: &loggingv1.Cloudwatch{
					Region:  "anumber1",
					GroupBy: loggingv1.LogGroupByNamespaceName,
				},
			},
			Secret: &loggingv1.OutputSecretSpec{
				Name: "my-secret",
			},
		}
		secrets = map[string]*corev1.Secret{
			output.Secret.Name: {
				Data: map[string][]byte{
					"aws_access_key_id":     nil,
					"aws_secret_access_key": nil,
				},
			},
		}
	)

	Context("for cloudwatch output ", func() {
		BeforeEach(func() {
			g = generator.MakeGenerator()
		})

		Context("grouped by log type", func() {
			BeforeEach(func() {
				output.Cloudwatch.GroupBy = loggingv1.LogGroupByLogType
			})
			It("should provide a valid configuration", func() {
				expConf := `
<label @MY_CLOUDWATCH>
 <filter kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_** journal.** system.var.log** var.log.pods.openshift-*_** var.log.pods.default_** var.log.pods.kube-*_**>
     @type record_modifier
    <record>
      cw_group_name infrastructure
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>
  
  <filter kubernetes.** var.log.pods.**>
    @type record_modifier
    <record>
      cw_group_name application
      cw_stream_name ${tag}
    </record>
  </filter>
  
  <filter linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
    @type record_modifier
    <record>
      cw_group_name audit
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>
  
  <match **>
    @type cloudwatch_logs
    auto_create_stream true
    region anumber1
    log_group_name_key cw_group_name
    log_stream_name_key cw_stream_name
    remove_log_stream_name_key true
    remove_log_group_name_key true
    concurrency 2
    aws_key_id "#{open('/var/run/ocp-collector/secrets/my-secret/aws_access_key_id','r') do |f|f.read.strip end}"
    aws_sec_key "#{open('/var/run/ocp-collector/secrets/my-secret/aws_secret_access_key','r') do |f|f.read.strip end}"
    include_time_key true
    log_rejected_request true
  </match>
</label>
`
				es := Conf(nil, secrets[output.Secret.Name], output, nil)
				results, err := g.GenerateConf(es...)
				Expect(err).To(BeNil())
				Expect(results).To(EqualTrimLines(expConf))
			})
		})
		Context("grouped by namespace", func() {
			BeforeEach(func() {
				output.Cloudwatch.GroupBy = loggingv1.LogGroupByNamespaceName
			})
			It("should provide a valid configuration", func() {
				expConf := `
<label @MY_CLOUDWATCH>
  <filter kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_** journal.** system.var.log** var.log.pods.openshift-*_** var.log.pods.default_** var.log.pods.kube-*_**>
    @type record_modifier
    <record>
      cw_group_name infrastructure
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>

  <filter kubernetes.** var.log.pods.**>
    @type record_modifier
    <record>
      cw_group_name ${record['kubernetes']['namespace_name']}
      cw_stream_name ${tag}
    </record>
  </filter>

  <filter linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
    @type record_modifier
    <record>
      cw_group_name audit
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>
  
  <match **>
    @type cloudwatch_logs
    auto_create_stream true
    region anumber1
    log_group_name_key cw_group_name
    log_stream_name_key cw_stream_name
    remove_log_stream_name_key true
    remove_log_group_name_key true
    concurrency 2
    aws_key_id "#{open('/var/run/ocp-collector/secrets/my-secret/aws_access_key_id','r') do |f|f.read.strip end}"
    aws_sec_key "#{open('/var/run/ocp-collector/secrets/my-secret/aws_secret_access_key','r') do |f|f.read.strip end}"
    include_time_key true
    log_rejected_request true
  </match>
</label>
`

				es := Conf(nil, secrets[output.Secret.Name], output, nil)
				results, err := g.GenerateConf(es...)
				Expect(err).To(BeNil())
				Expect(results).To(EqualTrimLines(expConf))
			})
		})
		Context("grouped by namespace UUID", func() {
			BeforeEach(func() {
				output.Cloudwatch.GroupBy = loggingv1.LogGroupByNamespaceUUID
				prefix := "foo"
				output.Cloudwatch.GroupPrefix = &prefix
			})
			It("should provide a valid configuration", func() {
				expConf := `
<label @MY_CLOUDWATCH>
  <filter kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_** journal.** system.var.log** var.log.pods.openshift-*_** var.log.pods.default_** var.log.pods.kube-*_**>
    @type record_modifier
    <record>
      cw_group_name foo.infrastructure
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>

  <filter kubernetes.** var.log.pods.**>
    @type record_modifier
    <record>
      cw_group_name foo.${record['kubernetes']['namespace_id']}
      cw_stream_name ${tag}
    </record>
  </filter>

  <filter linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
    @type record_modifier
    <record>
      cw_group_name foo.audit
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>
  
  <match **>
    @type cloudwatch_logs
    auto_create_stream true
    region anumber1
    log_group_name_key cw_group_name
    log_stream_name_key cw_stream_name
    remove_log_stream_name_key true
    remove_log_group_name_key true
    concurrency 2
    aws_key_id "#{open('/var/run/ocp-collector/secrets/my-secret/aws_access_key_id','r') do |f|f.read.strip end}"
    aws_sec_key "#{open('/var/run/ocp-collector/secrets/my-secret/aws_secret_access_key','r') do |f|f.read.strip end}"
    include_time_key true
    log_rejected_request true
  </match>
</label>
`

				es := Conf(nil, secrets[output.Secret.Name], output, nil)
				results, err := g.GenerateConf(es...)
				Expect(err).To(BeNil())
				Expect(results).To(EqualTrimLines(expConf))
			})
		})
	})
})

var _ = Describe("Generating fluentd config for sts", func() {
	var (
		g      generator.Generator
		output = loggingv1.OutputSpec{
			Type: loggingv1.OutputTypeCloudwatch,
			Name: "my-cloudwatch",
			OutputTypeSpec: loggingv1.OutputTypeSpec{
				Cloudwatch: &loggingv1.Cloudwatch{
					Region:  "anumber1",
					GroupBy: loggingv1.LogGroupByNamespaceName,
				},
			},
			Secret: &loggingv1.OutputSecretSpec{
				Name: "my-secret",
			},
		}
		roleArn = "arn:aws:iam::123456789012:role/my-role-to-assume"
		secrets = map[string]*corev1.Secret{
			output.Secret.Name: {
				Data: map[string][]byte{
					"role_arn": []byte(roleArn),
				},
			},
		}
		webIdentityTokenFile = path.Join(constants.AWSWebIdentityTokenMount, constants.AWSWebIdentityTokenFilePath)
	)

	Context("for cloudwatch sts output ", func() {
		BeforeEach(func() {
			g = generator.MakeGenerator()
		})

		Context("grouped by log type", func() {
			BeforeEach(func() {
				output.Cloudwatch.GroupBy = loggingv1.LogGroupByLogType
			})
			It("should provide a valid configuration for sts", func() {
				expConf := `
<label @MY_CLOUDWATCH>
 <filter kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_** journal.** system.var.log** var.log.pods.openshift-*_** var.log.pods.default_** var.log.pods.kube-*_**>
     @type record_modifier
    <record>
      cw_group_name infrastructure
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>
  
  <filter kubernetes.** var.log.pods.**>
    @type record_modifier
    <record>
      cw_group_name application
      cw_stream_name ${tag}
    </record>
  </filter>
  
  <filter linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
    @type record_modifier
    <record>
      cw_group_name audit
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>
  
  <match **>
    @type cloudwatch_logs
    auto_create_stream true
    region anumber1
    log_group_name_key cw_group_name
    log_stream_name_key cw_stream_name
    remove_log_stream_name_key true
    remove_log_group_name_key true
    concurrency 2
    <web_identity_credentials>
      role_arn "` + roleArn + `"
      web_identity_token_file "` + webIdentityTokenFile + `"
      role_session_name "` + constants.AWSRoleSessionName + `"
    </web_identity_credentials>    
    include_time_key true
    log_rejected_request true
  </match>
</label>
`
				es := Conf(nil, secrets[output.Secret.Name], output, nil)
				results, err := g.GenerateConf(es...)
				Expect(err).To(BeNil())
				Expect(results).To(EqualTrimLines(expConf))
			})
		})
		Context("grouped by namespace using sts", func() {
			BeforeEach(func() {
				output.Cloudwatch.GroupBy = loggingv1.LogGroupByNamespaceName
			})
			It("should provide a valid configuration for sts", func() {
				expConf := `
<label @MY_CLOUDWATCH>
  <filter kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_** journal.** system.var.log** var.log.pods.openshift-*_** var.log.pods.default_** var.log.pods.kube-*_**>
    @type record_modifier
    <record>
      cw_group_name infrastructure
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>

  <filter kubernetes.** var.log.pods.**>
    @type record_modifier
    <record>
      cw_group_name ${record['kubernetes']['namespace_name']}
      cw_stream_name ${tag}
    </record>
  </filter>

  <filter linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
    @type record_modifier
    <record>
      cw_group_name audit
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>
  
  <match **>
    @type cloudwatch_logs
    auto_create_stream true
    region anumber1
    log_group_name_key cw_group_name
    log_stream_name_key cw_stream_name
    remove_log_stream_name_key true
    remove_log_group_name_key true
    concurrency 2
    <web_identity_credentials>
      role_arn "` + roleArn + `"
      web_identity_token_file "` + webIdentityTokenFile + `"
      role_session_name "` + constants.AWSRoleSessionName + `"
    </web_identity_credentials>      
    include_time_key true
    log_rejected_request true
  </match>
</label>
`

				es := Conf(nil, secrets[output.Secret.Name], output, nil)
				results, err := g.GenerateConf(es...)
				Expect(err).To(BeNil())
				Expect(results).To(EqualTrimLines(expConf))
			})
		})
		Context("grouped by namespace UUID using sts", func() {
			BeforeEach(func() {
				output.Cloudwatch.GroupBy = loggingv1.LogGroupByNamespaceUUID
				prefix := "foo"
				output.Cloudwatch.GroupPrefix = &prefix
			})
			It("should provide a valid configuration for sts", func() {
				expConf := `
<label @MY_CLOUDWATCH>
  <filter kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_** journal.** system.var.log** var.log.pods.openshift-*_** var.log.pods.default_** var.log.pods.kube-*_**>
    @type record_modifier
    <record>
      cw_group_name foo.infrastructure
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>

  <filter kubernetes.** var.log.pods.**>
    @type record_modifier
    <record>
      cw_group_name foo.${record['kubernetes']['namespace_id']}
      cw_stream_name ${tag}
    </record>
  </filter>

  <filter linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
    @type record_modifier
    <record>
      cw_group_name foo.audit
      cw_stream_name ${record['hostname']}.${tag}
    </record>
  </filter>
  
  <match **>
    @type cloudwatch_logs
    auto_create_stream true
    region anumber1
    log_group_name_key cw_group_name
    log_stream_name_key cw_stream_name
    remove_log_stream_name_key true
    remove_log_group_name_key true
    concurrency 2
    <web_identity_credentials>
      role_arn "` + roleArn + `"
      web_identity_token_file "` + webIdentityTokenFile + `"
      role_session_name "` + constants.AWSRoleSessionName + `"
    </web_identity_credentials>     
    include_time_key true
    log_rejected_request true
  </match>
</label>
`

				es := Conf(nil, secrets[output.Secret.Name], output, nil)
				results, err := g.GenerateConf(es...)
				Expect(err).To(BeNil()) //
				Expect(results).To(EqualTrimLines(expConf))
			})
		})
	})
})

func TestFluendConfGenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fluend Conf Generation")
}
