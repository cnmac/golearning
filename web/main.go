package main

import (
	"fmt"
	"github.com/cnmac/golearning/web/bootstrap"
	"github.com/cnmac/golearning/web/middleware/identity"
	"github.com/cnmac/golearning/web/routes"
)

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	// init application
	app := bootstrap.New("Go web", "mc")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
