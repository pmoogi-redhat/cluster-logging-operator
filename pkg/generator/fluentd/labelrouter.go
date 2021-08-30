package fluentd

import (
	. "github.com/openshift/cluster-logging-operator/pkg/generator"
)

type LabelRouter struct {
	Routes []Element
}

func (lr LabelRouter) Name() string {
	return "labelRouterTemplate"
}

func (lr LabelRouter) Template() string {
	return `{{define "` + lr.Name() + `" -}}
@type label_router
{{compose .Routes}}
{{end}}`
}

type Route struct {
	RoutePipeline Element
}

func (r Route) Name() string {
	return "routeTemplate"
}

func (r Route) Template() string {
	return `{{define "` + r.Name() + `" -}}
<route>
{{compose_one .RoutePipeline| indent 2}}
</route>
{{end}}`
}

type RoutePipeline struct {
	Pipeline  string
	RouteData Element
}

func (p RoutePipeline) Name() string {
	return "routePipelineTemplate"
}

func (p RoutePipeline) Template() string {
	return `{{define "` + p.Name() + `" -}}
@label {{.Pipeline}}
<match>
{{compose_one .RouteData | indent 2}}
</match>
{{end}}`
}

type RouteData struct {
	// Labels is an array of "<key>:<value>" strings
	Labels     Element
	Namespaces Element
}

func (rd RouteData) Name() string {
	return "routeDataTemplate"
}

func (rd RouteData) Template() string {
	return `{{define "` + rd.Name() + `" -}}
{{kv .Namespaces -}}
{{kv .Labels -}}
{{end}}`
}
