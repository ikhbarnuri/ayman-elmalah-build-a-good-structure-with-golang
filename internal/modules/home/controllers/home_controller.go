package controllers

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	articleRepository repositories.ArticleRepositoryInterface
}

func New() *Controller {
	return &Controller{
		articleRepository: repositories.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	//html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	//	"title": "Home Page",
	//})
	c.JSON(http.StatusOK, gin.H{
		"articles": controller.articleRepository.List(10),
	})
}
