package routes

import (
	homeCtrl "ayman-elmalah-build-a-good-structure-with-golang/internal/modules/home/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)
}
