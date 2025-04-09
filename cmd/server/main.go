package main

import (
	"Spotify/internal/model"
	"Spotify/internal/routes"
	"github.com/gin-gonic/gin"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://admin:pass@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&model.Playlist{}, &model.Song{}, &model.Album{}, &model.PlaylistSongs{})
	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}

	r := gin.Default()
	routes.SetupRouter(db, r)

	r.Run(":8081")
}
