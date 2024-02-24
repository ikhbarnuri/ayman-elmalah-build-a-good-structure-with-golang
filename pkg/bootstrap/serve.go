package bootstrap

import (
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/config"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/database"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/html"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/routing"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()
	router := routing.GetRouter()

	static.LoadStatic(router)

	html.LoadHTML(router)

	routing.RegisterRoutes()

	routing.Serve()
}
