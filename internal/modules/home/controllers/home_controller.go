package controllers

import (
	"ayman-elmalah-build-a-good-structure-with-golang/package/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title": "Home Page",
	})
}
