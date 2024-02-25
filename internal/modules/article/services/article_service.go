package services

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/models"
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/repositories"
)

type ArticleService struct {
	articleREspository repositories.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleREspository: repositories.New(),
	}
}

func (articleService ArticleService) GetFeaturedArticles() []models.Article {
	return articleService.articleREspository.List(4)
}

func (articleService ArticleService) GetStoriesArticles() []models.Article {
	return articleService.articleREspository.List(4)
}
