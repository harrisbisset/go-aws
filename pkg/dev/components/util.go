package components

import (
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/dev/render/views/view_error"
)

var DefaultBuildError = fnet.NewError("build error", view_error.RawBuildError()).
	Error("failed at build stage").
	Build()
