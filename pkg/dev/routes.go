package dev

import (
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/dev/components/pages"
)

func (c cfg) createRoutes() {
	fnet.HandleComponent(fnet.GET, "/", pages.ShowIndex)
	fnet.HandleComponent(fnet.GET, "/instances", pages.ShowInstances)

}
