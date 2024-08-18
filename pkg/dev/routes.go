package dev

import "github.com/harrisbisset/fnet"

func (c cfg) createRoutes() {
	fnet.HandleComponent(fnet.GET, "/")
}
