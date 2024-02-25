package seeder

import (
	models2 "ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/models"
	"ayman-elmalah-build-a-good-structure-with-golang/internal/modules/user/models"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/database"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Seed() {
	db := database.Connection()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)
	if err != nil {
		log.Fatal("hash password error")
		return
	}

	user := models.User{
		Name:     "Jinzhu",
		Email:    "jinzhu@mail.com",
		Password: string(hashPassword),
	}
	db.Create(&user)

	log.Printf("User created successfully with email address %s \n", user.Email)

	for i := 0; i <= 10; i++ {
		article := models2.Article{
			Title:   fmt.Sprintf("Randowm Title %v", i),
			Content: fmt.Sprintf("Lorem ipsim dolor si amet %v", i),
			UserID:  user.ID,
			User:    user,
		}
		db.Create(&article)

		log.Printf("Aticle created successfully with title %s\n", article.Title)
	}

	log.Println("Seeder done")
}
