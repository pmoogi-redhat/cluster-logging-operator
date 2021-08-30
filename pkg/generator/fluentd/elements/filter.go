package elements

import (
	. "github.com/openshift/cluster-logging-operator/pkg/generator"
)

type Filter struct {
	Desc      string
	MatchTags string
	Element   Element
}

func (f Filter) Name() string {
	return "filterTemplate"
}

func (f Filter) Template() string {
	return `{{define "` + f.Name() + `" -}}
{{if .Desc -}}
#{{.Desc}}
{{end -}}
<filter {{.MatchTags}}>
{{compose_one .Element | indent 2}}
</filter>
{{end}}
`
}
