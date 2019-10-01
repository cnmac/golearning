package main

import (
	"fmt"
	"github.com/cnmac/golearning/lottery/bootstrap"
	"github.com/cnmac/golearning/lottery/middleware/identity"
	"github.com/cnmac/golearning/lottery/routes"
)

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	// init application
	app := bootstrap.New("Go lottery", "mc")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
