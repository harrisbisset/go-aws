package dev

import "github.com/harrisbisset/fnet"

func StartDev(addr string, cidr string) {
	if addr == "" {
		addr = "0.0.0.0"
	}
	if cidr == "" {
		cidr = "80"
	}

	fnet.Start(addr, cidr)
}
