package pages

import "github.com/harrisbisset/fnet"

var LoginPage = fnet.NewComponent("Login Page").
	View().
	Error().
	Build()
