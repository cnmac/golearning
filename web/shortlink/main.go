package main

import "github.com/cnmac/golearning/web/shortlink/app"

func main() {
	a := app.App{}
	a.Initialize()
	a.Run(":8000")
}
