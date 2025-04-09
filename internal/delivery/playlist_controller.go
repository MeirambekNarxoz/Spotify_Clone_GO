package delivery

import (
	"Spotify/internal/model"
	"Spotify/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PlaylistController структура для работы с плейлистами
type PlaylistController struct {
	playlistService service.PlaylistService
}

// NewPlaylistController создает новый контроллер для плейлистов
func NewPlaylistController(playlistService service.PlaylistService) *PlaylistController {
	return &PlaylistController{
		playlistService: playlistService,
	}
}

// CreatePlaylist создает новый плейлист
func (p *PlaylistController) CreatePlaylist(c *gin.Context) {
	var playlist model.Playlist
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := p.playlistService.CreatePlaylist(&playlist); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Playlist created successfully"})
}

// GetPlaylistByID возвращает плейлист по ID
func (p *PlaylistController) GetPlaylistByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	playlist, err := p.playlistService.GetPlaylistByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	c.JSON(http.StatusOK, playlist)
}

// GetAllPlaylists возвращает все плейлисты
func (p *PlaylistController) GetAllPlaylists(c *gin.Context) {
	playlists, err := p.playlistService.GetAllPlaylists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve playlists"})
		return
	}

	c.JSON(http.StatusOK, playlists)
}

// GetPlaylistsByUserID возвращает все плейлисты для пользователя
func (p *PlaylistController) GetPlaylistsByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	playlists, err := p.playlistService.GetPlaylistsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve playlists"})
		return
	}

	c.JSON(http.StatusOK, playlists)
}

// UpdatePlaylist обновляет плейлист
func (p *PlaylistController) UpdatePlaylist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var playlist model.Playlist
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	playlist.ID = uint(id)
	if err := p.playlistService.UpdatePlaylist(&playlist); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist updated successfully"})
}

// DeletePlaylist удаляет плейлист
func (p *PlaylistController) DeletePlaylist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := p.playlistService.DeletePlaylist(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist deleted successfully"})
}
