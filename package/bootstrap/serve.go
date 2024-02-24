package bootstrap

import (
	"ayman-elmalah-build-a-good-structure-with-golang/package/config"
	"ayman-elmalah-build-a-good-structure-with-golang/package/html"
	"ayman-elmalah-build-a-good-structure-with-golang/package/routing"
	"ayman-elmalah-build-a-good-structure-with-golang/package/static"
)

func Serve() {
	config.Set()

	routing.Init()
	router := routing.GetRouter()

	static.LoadStatic(router)

	html.LoadHTML(router)

	routing.RegisterRoutes()

	routing.Serve()
}
