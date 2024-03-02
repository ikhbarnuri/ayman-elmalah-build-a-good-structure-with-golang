package controllers

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/services"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	articleService services.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: services.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	article, err := controller.articleService.Find(id)
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Page not found	", "message": err.Error()})
		return
	}

	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Show article	", "article": article})
}
