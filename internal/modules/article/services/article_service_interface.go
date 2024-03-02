package services

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() responses.Articles
	GetStoriesArticles() responses.Articles
	Find(id int) (responses.Article, error)
}
