package migration

import (
	articleModel "ayman-elmalah-build-a-good-structure-with-golang/internal/modules/article/models"
	userModel "ayman-elmalah-build-a-good-structure-with-golang/internal/modules/user/models"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(userModel.User{}, articleModel.Article{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration done")
}
