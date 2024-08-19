package pages

import (
	"net/http"

	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_instances"
)

var InstancePage = fnet.NewComponent("Login Page").
	View(view_instances.Show()).
	Error(view_instances.Show()).
	Build()

func ShowLogin(w http.ResponseWriter, req *http.Request) {
	InstancePage.Render(w, req)
}
