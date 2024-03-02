package repositories

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/models"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []models.Article {
	var articles []models.Article

	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}

func (articleRepository *ArticleRepository) Find(id int) models.Article {
	var article models.Article

	articleRepository.DB.Joins("User").First(&article, id)

	return article
}
