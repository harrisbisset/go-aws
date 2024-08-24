package raw

import (
	"net/http"

	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/api"
	"github.com/harrisbisset/go-aws/pkg/dev/components"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_error"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_instances"
)

var IndexInstances = fnet.NewComponent("Index Instances").
	Error(0, components.DefaultBuildError).
	Error(1, fnet.NewError("api error", view_error.RawBuildError()).
		Error("api responded poorly").
		Build()).
	Build()

func ShowIndexInstances(w http.ResponseWriter, req *http.Request) {
	instances, err := api.GetInstanceList()
	if err != nil {
		IndexInstances.RenderError(1, w, req)
	}
	IndexInstances.View(view_instances.ListInstances(instances))
	IndexInstances.Render(w, req)
}

var InstanceConnection = fnet.NewComponent("Check Instance Connection").
	Error(0, components.DefaultBuildError).
	Error(1, fnet.NewError("api error", view_error.RawBuildError()).
		Error("api responded poorly").
		Build()).
	Build()

func ShowCheckInstanceConnection(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	report, err := api.CheckInstanceConnection(id)
	if err != nil {
		InstanceConnection.RenderError(1, w, req)
	}

	InstanceConnection.View(view_instances.InstanceReport(report))
	InstanceConnection.Render(w, req)
}
