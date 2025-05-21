package delivery

import (
	"Spotify/internal/model"
	"Spotify/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AlbumController структура для работы с альбомами
type AlbumController struct {
	albumService service.AlbumService
}

// NewAlbumController создает новый контроллер для альбомов
func NewAlbumController(albumService service.AlbumService) *AlbumController {
	return &AlbumController{
		albumService: albumService,
	}
}

// CreateAlbum создает новый альбом
func (a *AlbumController) CreateAlbum(c *gin.Context) {
	var album model.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := a.albumService.CreateAlbum(&album); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Album created successfully"})
}

// GetAlbumByID возвращает альбом по ID
// GetAlbumByID godoc
// @Security BearerToken
// @Summary Get album by ID
// @Security ApiKeyAuth
// @Description Get an album by its ID (public)
// @Tags albums
// @Produce json
// @Param id path int true "Album ID"
// @Success 200 {object} model.Album
// @Failure 404 {object} ErrorResponse
// @Router /albums/{id} [get]
func (a *AlbumController) GetAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	album, err := a.albumService.GetAlbumByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

// GetAllAlbums возвращает все альбомы
// GetAllAlbums godoc
// @Summary Get all albums
// @Description Get list of all albums (public)
// @Tags albums
// @Security BearerToken

// @Produce json
// @Success 200 {array} model.Album
// @Router /albums/ [get]
func (a *AlbumController) GetAllAlbums(c *gin.Context) {
	albums, err := a.albumService.GetAllAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve albums"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

// UpdateAlbum обновляет альбом
func (a *AlbumController) UpdateAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var album model.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	album.ID = uint(id)
	if err := a.albumService.UpdateAlbum(&album); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album updated successfully"})
}

// DeleteAlbum удаляет альбом
func (a *AlbumController) DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := a.albumService.DeleteAlbum(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}
