package cmd

import (
	"ayman-elmalah-build-a-good-structure-with-golang/package/config"
	"ayman-elmalah-build-a-good-structure-with-golang/package/routing"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  "Application will be served on host and port defined in config",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	config.Set()
	configs := config.Get()

	routing.Init()
	router := routing.GetRouter()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "pong",
			"app_name": configs.App.Name,
		})
	})

	routing.Serve()
}
