package pages

import (
	"net/http"

	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/dev/components"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_instances"
)

var InstancePage = fnet.NewComponent("Instance Page").
	View(view_instances.Show()).
	Error(0, components.DefaultBuildError).
	Build()

func ShowInstances(w http.ResponseWriter, req *http.Request) {
	InstancePage.Render(w, req)
}
