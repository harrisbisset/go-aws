package raw

import (
	"fmt"
	"net/http"

	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/api"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_error"
)

var IndexInstances = fnet.NewComponent("Index Instances").
	Error(0, fnet.RespError("build error", view_error.RawBuildError())).
	Error(1, fnet.RespError("api error", view_error.RawBuildError())).
	Build()

func Show(w http.ResponseWriter, req *http.Request) {
	instanceList, err := api.GetInstanceList()
	if err != nil {
		IndexInstances.RenderError(1, w, req)
	}
	c := instanceList.Reservations
	fmt.Println(c)
	// IndexInstances.View(view_instances.ListInstances())
	IndexInstances.Render(w, req)
}
