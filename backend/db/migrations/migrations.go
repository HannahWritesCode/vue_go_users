package migrations

import (
	"vue-go-users.com/db"
	"vue-go-users.com/models"
)

// Migrate runs the migrations
func Migrate() {
	// Migrate DB
	db.GetDB().AutoMigrate(
		&models.User{},
	)
}
