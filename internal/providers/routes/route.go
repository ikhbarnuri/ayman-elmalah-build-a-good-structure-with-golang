package routes

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {
	routes.Routes(router)
}
