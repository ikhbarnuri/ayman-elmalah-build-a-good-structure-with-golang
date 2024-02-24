package routes

import (
	"ayman-elmalah-build-a-good-structure-with-golang/package/config"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	configs := config.Get()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "pong",
			"app_name": configs.App.Name,
		})
	})
}
