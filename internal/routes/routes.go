package routes

import (
	"Spotify/internal/delivery"
	"Spotify/internal/repository"
	"Spotify/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, router *gin.Engine) *gin.Engine {
	// Репозитории
	albumRepo := repository.NewAlbumRepository(db)
	songRepo := repository.NewSongRepository(db)
	playlistRepo := repository.NewPlaylistRepository(db)           // ✅ новый
	playlistSongsRepo := repository.NewPlaylistSongsRepository(db) // Добавляем новый репозиторий для связи песен с плейлистами

	// Сервисы
	albumService := service.NewAlbumService(albumRepo)
	songService := service.NewSongService(songRepo)
	playlistService := service.NewPlaylistService(playlistRepo)                // ✅ новый
	playlistSongsService := service.NewPlaylistSongsService(playlistSongsRepo) // Добавляем новый сервис

	// Контроллеры
	albumController := delivery.NewAlbumController(albumService)
	songController := delivery.NewSongController(songService)
	playlistController := delivery.NewPlaylistController(playlistService)                // ✅ новый
	playlistSongsController := delivery.NewPlaylistSongsController(playlistSongsService) // Добавляем новый контроллер

	// Альбомы
	albumsGroup := router.Group("/albums")
	{
		albumsGroup.GET("/", albumController.GetAllAlbums)
		albumsGroup.GET("/:id", albumController.GetAlbumByID)
		albumsGroup.POST("/", albumController.CreateAlbum)
		albumsGroup.PUT("/:id", albumController.UpdateAlbum)
		albumsGroup.DELETE("/:id", albumController.DeleteAlbum)
	}

	// Песни
	songGroup := router.Group("/songs")
	{
		songGroup.GET("/", songController.GetAllSongs)
		songGroup.GET("/:id", songController.GetSongByID)
		songGroup.POST("/", songController.CreateSong)
		songGroup.PUT("/:id", songController.UpdateSong)
		songGroup.DELETE("/:id", songController.DeleteSong)
	}

	// Плейлисты
	playlistGroup := router.Group("/playlists")
	{
		playlistGroup.GET("/", playlistController.GetAllPlaylists)
		playlistGroup.GET("/:id", playlistController.GetPlaylistByID)
		playlistGroup.POST("/", playlistController.CreatePlaylist)
		playlistGroup.PUT("/:id", playlistController.UpdatePlaylist)
		playlistGroup.DELETE("/:id", playlistController.DeletePlaylist)
	}

	// Песни в плейлистах
	playlistSongsGroup := router.Group("/playlists/:id/songs")
	{
		playlistSongsGroup.POST("/", playlistSongsController.AddSongToPlaylist)           // Добавление песни в плейлист
		playlistSongsGroup.DELETE("/:id", playlistSongsController.RemoveSongFromPlaylist) // Удаление песни из плейлиста
		playlistSongsGroup.GET("/", playlistSongsController.GetSongsInPlaylist)           // Получение всех песен в плейлисте
	}

	return router
}
