package delivery

import (
	"Spotify/internal/model"
	"Spotify/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SongController структура для работы с песнями
type SongController struct {
	songService service.SongService
}

// NewSongController создает новый контроллер для песен
func NewSongController(songService service.SongService) *SongController {
	return &SongController{
		songService: songService,
	}
}

// GetAllSongs
// @Security BearerToken

// @Summary Get all songs
// @Tags songs
// @Produce json
// @Success 200 {array} model.Song
// @Router /songs [get]
func (s *SongController) GetAllSongs(c *gin.Context) {
	songs, err := s.songService.GetAllSongs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve songs"})
		return
	}
	c.JSON(http.StatusOK, songs)
}

func (s *SongController) CreateSong(c *gin.Context) {
	var song model.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := s.songService.CreateSong(&song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Song created successfully"})
}

// GetSongByID возвращает песню по ID
// GetSongByID godoc
// @Security BearerToken

// @Summary Get song by ID
// @Description Get a song by its ID (public)
// @Tags songs
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} model.Song
// @Failure 404 {object} ErrorResponse
// @Router /songs/{id} [get]
func (s *SongController) GetSongByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	song, err := s.songService.GetSongByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	c.JSON(http.StatusOK, song)
}

// GetSongsByAlbumID возвращает все песни по albumID
func (s *SongController) GetSongsByAlbumID(c *gin.Context) {
	albumID, err := strconv.Atoi(c.Param("albumID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Album ID"})
		return
	}

	songs, err := s.songService.GetSongsByAlbumID(uint(albumID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve songs"})
		return
	}

	c.JSON(http.StatusOK, songs)
}

// UpdateSong обновляет песню
func (s *SongController) UpdateSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var song model.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	song.ID = uint(id)
	if err := s.songService.UpdateSong(&song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song updated successfully"})
}

// DeleteSong удаляет песню
func (s *SongController) DeleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := s.songService.DeleteSong(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully"})
}
