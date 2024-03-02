package routes

import (
	articleRoutes "ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/routes"
	homeRoutes "ayman-elmalah-build-a-good-structure-with-golang/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
}
