package bootstrap

import (
	"ayman-elmalah-build-a-good-structure-with-golang/package/config"
	"ayman-elmalah-build-a-good-structure-with-golang/package/html"
	"ayman-elmalah-build-a-good-structure-with-golang/package/routing"
)

func Serve() {
	config.Set()

	routing.Init()
	html.LoadHTML(routing.GetRouter())
	routing.RegisterRoutes()
	routing.Serve()
}
