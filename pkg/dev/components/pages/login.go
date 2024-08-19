package pages

import (
	"net/http"

	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_login"
)

var LoginPage = fnet.NewComponent("Login Page").
	View(view_login.Show()).
	Error(view_login.Show()).
	Build()

func ShowLogin(w http.ResponseWriter, req *http.Request) {
	LoginPage.Render(w, req)
}
