package controllers

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/services"
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
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error converting the id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hello article"})
}
