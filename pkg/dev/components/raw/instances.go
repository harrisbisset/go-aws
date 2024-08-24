package raw

import (
	"net/http"

	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/api"
	"github.com/harrisbisset/go-aws/pkg/dev/components"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_error"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_instances"
)

var InstanceConnection = fnet.NewComponent("Instance Connection").
	Error(0, components.DefaultBuildError).
	Error(1, fnet.NewError("api error", view_error.RawBuildError()).
		Error("api responded poorly").
		Build()).
	Build()

func ShowInstanceConnection(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	zone := req.PathValue("zone")
	report, err := api.CheckInstanceConnection(&id, &zone)
	if err != nil {
		InstanceConnection.RenderError(1, w, req)
	}

	InstanceConnection.View(view_instances.InstanceReport(report))
	InstanceConnection.Render(w, req)
}
