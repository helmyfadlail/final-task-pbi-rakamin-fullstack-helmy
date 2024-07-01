package database

import (
	"image-profile-api/app"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(&app.User{}, &app.Photo{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}