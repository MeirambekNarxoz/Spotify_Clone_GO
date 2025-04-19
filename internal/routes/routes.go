package routes

import (
	"Spotify/internal/delivery"
	"Spotify/internal/middleware"
	"Spotify/internal/repository"
	"Spotify/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, router *gin.Engine) *gin.Engine {
	// Репозитории
	albumRepo := repository.NewAlbumRepository(db)
	songRepo := repository.NewSongRepository(db)
	playlistRepo := repository.NewPlaylistRepository(db)
	playlistSongsRepo := repository.NewPlaylistSongsRepository(db)

	// Сервисы
	albumService := service.NewAlbumService(albumRepo)
	songService := service.NewSongService(songRepo)
	playlistService := service.NewPlaylistService(playlistRepo)
	playlistSongsService := service.NewPlaylistSongsService(playlistSongsRepo)
	JWTService := service.NewJWTService()

	// Контроллеры
	albumController := delivery.NewAlbumController(albumService)
	songController := delivery.NewSongController(songService)
	playlistController := delivery.NewPlaylistController(playlistService)
	playlistSongsController := delivery.NewPlaylistSongsController(playlistSongsService)

	// Альбомы
	albumsGroup := router.Group("/albums", middleware.AuthMiddleware(JWTService.SecretKey), middleware.RequireModerator())
	{
		albumsGroup.GET("/", albumController.GetAllAlbums)
		albumsGroup.GET("/:id", albumController.GetAlbumByID)
		albumsGroup.POST("/", albumController.CreateAlbum)
		albumsGroup.PUT("/:id", albumController.UpdateAlbum)
		albumsGroup.DELETE("/:id", albumController.DeleteAlbum)
	}

	router.GET("/songs/", songController.GetAllSongs)
	router.GET("/songs/:id", songController.GetSongByID)
	// Песни
	songGroup := router.Group("/songs", middleware.AuthMiddleware(JWTService.SecretKey))
	{
		songGroup.POST("/", songController.CreateSong)
		songGroup.PUT("/:id", songController.UpdateSong)
		songGroup.DELETE("/:id", songController.DeleteSong)
	}

	// Плейлисты
	playlistGroup := router.Group("/playlists", middleware.AuthMiddleware(JWTService.SecretKey))
	{
		playlistGroup.GET("/", playlistController.GetAllPlaylists)
		playlistGroup.GET("/:id", playlistController.GetPlaylistByID)
		playlistGroup.POST("/", playlistController.CreatePlaylist)
		playlistGroup.PUT("/:id", playlistController.UpdatePlaylist)
		playlistGroup.DELETE("/:id", playlistController.DeletePlaylist)
	}

	// Песни в плейлистах
	playlistSongsGroup := router.Group("/playlists/:id/songs", middleware.AuthMiddleware(JWTService.SecretKey))
	{
		playlistSongsGroup.POST("/", playlistSongsController.AddSongToPlaylist)
		playlistSongsGroup.DELETE("/:songId", playlistSongsController.RemoveSongFromPlaylist)
		playlistSongsGroup.GET("/", playlistSongsController.GetSongsInPlaylist)
	}

	return router
}
