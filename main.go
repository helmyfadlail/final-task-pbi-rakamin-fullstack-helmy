package main

import (
	"image-profile-api/database"
	"image-profile-api/router"
	"log"
)

func main() {
	database.InitDB()
	database.Migrate()

	r := router.InitRouter()

	log.Println("Server started on port 8080")
	r.Run(":8080")
}
