package cmd

import (
	"ayman-elmalah-build-a-good-structure-with-golang/package/config"
	"fmt"
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "pong",
			"app_name": configs.App.Name,
		})
	})

	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}
