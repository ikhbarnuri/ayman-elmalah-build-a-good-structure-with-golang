package routing

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/providers/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	routes.RegisterRoute(GetRouter())
}
