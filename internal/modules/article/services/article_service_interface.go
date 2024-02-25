package services

import "ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticles() []models.Article
	GetStoriesArticles() []models.Article
}
