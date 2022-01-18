package main

import (
	"gwftemplate/data"
	"gwftemplate/handlers"
	"gwftemplate/middleware"

	"github.com/ck3g/gwf"
)

type application struct {
	App        *gwf.GWF
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
