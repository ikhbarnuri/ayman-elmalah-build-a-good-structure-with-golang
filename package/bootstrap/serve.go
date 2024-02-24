package bootstrap

import (
	"ayman-elmalah-build-a-good-structure-with-golang/package/config"
	"ayman-elmalah-build-a-good-structure-with-golang/package/routing"
)

func Serve() {
	config.Set()

	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
