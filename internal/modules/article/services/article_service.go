package services

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/repositories"
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/responses"
	"errors"
)

type ArticleService struct {
	articleRespository repositories.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRespository: repositories.New(),
	}
}

func (articleService ArticleService) GetFeaturedArticles() responses.Articles {
	articles := articleService.articleRespository.List(4)

	return responses.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() responses.Articles {
	articles := articleService.articleRespository.List(4)

	return responses.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (responses.Article, error) {
	var response responses.Article

	article := articleService.articleRespository.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")

	}

	return responses.ToArticle(article), nil
}
