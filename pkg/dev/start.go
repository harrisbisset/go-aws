package dev

import (
	"net/http"

	"github.com/harrisbisset/fnet"
)

func StartDev(addr string, cidr string) {
	cfg := createConfig()
	cfg.createRoutes()
	fnet.Start(addr, cidr, http.NewServeMux())
}
