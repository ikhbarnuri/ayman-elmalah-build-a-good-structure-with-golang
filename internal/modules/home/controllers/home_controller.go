package controllers

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	articleService services.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: services.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	//html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	//	"title": "Home Page",
	//})
	c.JSON(http.StatusOK, gin.H{
		"articles": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})
}
