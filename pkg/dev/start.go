package dev

import (
	"log"

	"github.com/harrisbisset/fnet"
)

func StartDev(addr string, cidr string) {
	cfg := createConfig()
	cfg.createRoutes()
	log.Print("starting server")
	fnet.Start(addr, cidr, nil)
}
