package delivery

import (
	"Spotify/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PlaylistSongsController struct {
	playlistSongsService service.PlaylistSongsService
}

func NewPlaylistSongsController(playlistSongsService service.PlaylistSongsService) *PlaylistSongsController {
	return &PlaylistSongsController{playlistSongsService: playlistSongsService}
}

// Добавить песню в плейлист
type AddSongRequest struct {
	SongID uint `json:"song_id"`
}

func (p *PlaylistSongsController) AddSongToPlaylist(c *gin.Context) {
	playlistID := c.Param("id")

	var req AddSongRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	// Преобразование параметров в uint
	playlistIDUint, err := strconv.ParseUint(playlistID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// Добавление песни в плейлист
	err = p.playlistSongsService.AddSongToPlaylist(uint(playlistIDUint), req.SongID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song added to playlist"})
}

// Удалить песню из плейлиста
func (p *PlaylistSongsController) RemoveSongFromPlaylist(c *gin.Context) {
	playlistID := c.Param("id")
	songID := c.Param("id")

	// Преобразование параметров в uint
	playlistIDUint, err := strconv.ParseUint(playlistID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}
	songIDUint, err := strconv.ParseUint(songID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// Удаление песни из плейлиста
	err = p.playlistSongsService.RemoveSongFromPlaylist(uint(playlistIDUint), uint(songIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song removed from playlist"})
}

// Получить все песни в плейлисте
func (p *PlaylistSongsController) GetSongsInPlaylist(c *gin.Context) {
	playlistID := c.Param("id")

	// Преобразование параметра в uint
	playlistIDUint, err := strconv.ParseUint(playlistID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// Получение песен в плейлисте
	songs, err := p.playlistSongsService.GetSongsInPlaylist(uint(playlistIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"songs": songs})
}
