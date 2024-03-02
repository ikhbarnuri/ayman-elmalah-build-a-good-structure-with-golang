package routes

import (
	articleCtrl "ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articleController := articleCtrl.New()
	router.GET("/articles/:id", articleController.Show)
}
