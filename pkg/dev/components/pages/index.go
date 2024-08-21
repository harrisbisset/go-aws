package pages

import (
	"net/http"

	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_error"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_index"
)

var IndexPage = fnet.NewComponent("Index Page").
	View(view_index.Show()).
	Error(view_error.RawBuildError()).
	Build()

func ShowIndex(w http.ResponseWriter, req *http.Request) {
	IndexPage.Render(w, req)
}
