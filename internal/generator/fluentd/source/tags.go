package source

const (
	ApplicationTags               = "kubernetes.**"
	ApplicationTagsForMultilineEx = ApplicationTags + " var.log.pods.**"
	JournalTags                   = "journal.** system.var.log**"
	InfraContainerTags            = "kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_**"
	InfraTags                     = InfraContainerTags + " " + JournalTags
	InfraTagsForMultilineEx       = InfraTags + " var.log.pods.openshift-*_** var.log.pods.default_** var.log.pods.kube-*_**"
	AuditTags                     = "linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**"
)
