package bootstrap

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/database/migration"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/config"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
