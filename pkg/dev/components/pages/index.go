package pages

import (
	"net/http"

	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/dev/components"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_index"
)

var IndexPage = fnet.NewComponent("Index Page").
	View(view_index.Show()).
	Error(0, components.DefaultBuildError).
	Build()

func ShowIndex(w http.ResponseWriter, req *http.Request) {
	IndexPage.Render(w, req)
}
