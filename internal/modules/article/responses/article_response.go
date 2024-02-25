package responses

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/models"
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/user/responses"
	"fmt"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      responses.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article models.Article) Article {
	return Article{
		ID:        article.ID,
		Image:     "/assets/img/demopic/10.jpg",
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      responses.ToUser(article.User),
	}
}

func ToArticles(articles []models.Article) Articles {
	var response Articles

	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}

	return response
}
