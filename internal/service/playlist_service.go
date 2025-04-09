package service

import (
	"Spotify/internal/model"
	"Spotify/internal/repository"
	"errors"
)

// PlaylistService интерфейс для работы с плейлистами
type PlaylistService interface {
	CreatePlaylist(playlist *model.Playlist) error
	GetPlaylistByID(id uint) (*model.Playlist, error)
	GetPlaylistsByUserID(userID uint) ([]model.Playlist, error)
	GetAllPlaylists() ([]model.Playlist, error) // ➕ новый метод
	UpdatePlaylist(playlist *model.Playlist) error
	DeletePlaylist(id uint) error
}

type playlistService struct {
	playlistRepo repository.PlaylistRepository
}

// NewPlaylistService создает новый сервис для плейлистов
func NewPlaylistService(playlistRepo repository.PlaylistRepository) PlaylistService {
	return &playlistService{
		playlistRepo: playlistRepo,
	}
}

// GetAllPlaylists возвращает все плейлисты
func (s *playlistService) GetAllPlaylists() ([]model.Playlist, error) {
	return s.playlistRepo.GetAll()
}

// CreatePlaylist создает новый плейлист
func (s *playlistService) CreatePlaylist(playlist *model.Playlist) error {
	// Бизнес-логика для создания плейлиста
	if playlist.Name == "" {
		return errors.New("playlist name cannot be empty")
	}
	return s.playlistRepo.Create(playlist)
}

// GetPlaylistByID возвращает плейлист по ID
func (s *playlistService) GetPlaylistByID(id uint) (*model.Playlist, error) {
	return s.playlistRepo.GetByID(id)
}

// GetPlaylistsByUserID возвращает все плейлисты для пользователя
func (s *playlistService) GetPlaylistsByUserID(userID uint) ([]model.Playlist, error) {
	return s.playlistRepo.GetByUserID(userID)
}

// UpdatePlaylist обновляет плейлист
func (s *playlistService) UpdatePlaylist(playlist *model.Playlist) error {
	// Бизнес-логика для обновления плейлиста
	if playlist.Name == "" {
		return errors.New("playlist name cannot be empty")
	}
	return s.playlistRepo.Update(playlist)
}

// DeletePlaylist удаляет плейлист
func (s *playlistService) DeletePlaylist(id uint) error {
	return s.playlistRepo.Delete(id)
}
