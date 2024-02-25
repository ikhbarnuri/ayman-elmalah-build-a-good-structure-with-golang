package bootstrap

import (
	"ayman-elmalah-build-a-good-structure-with-golang/internal/database/seeder"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/config"
	"ayman-elmalah-build-a-good-structure-with-golang/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
