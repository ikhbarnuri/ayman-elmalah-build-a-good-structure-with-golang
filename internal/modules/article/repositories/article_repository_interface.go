package repositories

import models2 "ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []models2.Article
}
