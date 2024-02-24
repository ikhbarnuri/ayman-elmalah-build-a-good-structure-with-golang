package routing

import (
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/config"
	"fmt"
	"log"
)

func Serve() {
	configs := config.Get()

	router := GetRouter()
	err := router.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
	if err != nil {
		log.Fatal("Error in routing")
		return
	}
}
