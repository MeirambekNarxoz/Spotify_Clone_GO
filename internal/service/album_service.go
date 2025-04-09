package service

import (
	"Spotify/internal/model"
	"Spotify/internal/repository"
	"errors"
)

// AlbumService интерфейс для работы с альбомами
type AlbumService interface {
	CreateAlbum(album *model.Album) error
	GetAlbumByID(id uint) (*model.Album, error)
	GetAllAlbums() ([]model.Album, error)
	UpdateAlbum(album *model.Album) error
	DeleteAlbum(id uint) error
}

type albumService struct {
	albumRepo repository.AlbumRepository
}

// NewAlbumService создает новый сервис для альбомов
func NewAlbumService(albumRepo repository.AlbumRepository) AlbumService {
	return &albumService{
		albumRepo: albumRepo,
	}
}

// CreateAlbum создает новый альбом
func (s *albumService) CreateAlbum(album *model.Album) error {
	// Бизнес-логика для создания альбома
	if album.Name == "" {
		return errors.New("album name cannot be empty")
	}
	return s.albumRepo.Create(album)
}

// GetAlbumByID возвращает альбом по ID
func (s *albumService) GetAlbumByID(id uint) (*model.Album, error) {
	return s.albumRepo.GetByID(id)
}

// GetAllAlbums возвращает все альбомы
func (s *albumService) GetAllAlbums() ([]model.Album, error) {
	return s.albumRepo.GetAll()
}

// UpdateAlbum обновляет альбом
func (s *albumService) UpdateAlbum(album *model.Album) error {
	// Бизнес-логика для обновления альбома
	if album.Name == "" {
		return errors.New("album name cannot be empty")
	}
	return s.albumRepo.Update(album)
}

// DeleteAlbum удаляет альбом
func (s *albumService) DeleteAlbum(id uint) error {
	return s.albumRepo.Delete(id)
}
