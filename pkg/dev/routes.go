package dev

import (
	"log"
	"net/http"

	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/go-aws/pkg/dev/components/pages"
	"github.com/harrisbisset/go-aws/pkg/dev/components/raw"
)

func (c cfg) createRoutes() {
	fnet.HandleComponent(fnet.GET, "/", pages.ShowIndex)
	fnet.HandleComponent(fnet.GET, "/instance", pages.ShowInstances)
	fnet.HandleComponent(fnet.GET, "/instance/list", raw.ShowIndexInstances)
	fnet.HandleComponent(fnet.GET, "/instance/checkConnection/index/{id}/{zone}", raw.ShowIndexInstanceConnection)
	fnet.HandleComponent(fnet.GET, "/instance/checkConnection/{id}/{user}/{zone}", raw.ShowInstanceConnection)

	//static routing
	fnet.Handle(fnet.GET, "/dist/", http.StripPrefix("/dist/", fnet.NeuterFileSystem(http.FileServer(http.Dir("./pkg/dev/render/dist")))))
	fnet.Handle(fnet.GET, "/public/", http.StripPrefix("/public/", fnet.NeuterFileSystem(http.FileServer(http.Dir("./pkg/dev/render/public")))))
	log.Print("created routes")
}
