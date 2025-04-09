package service

import (
	"Spotify/internal/model"
	"Spotify/internal/repository"
	"errors"
)

// SongService интерфейс для работы с песнями
type SongService interface {
	CreateSong(song *model.Song) error
	GetSongByID(id uint) (*model.Song, error)
	GetSongsByAlbumID(albumID uint) ([]model.Song, error)
	GetAllSongs() ([]model.Song, error) // <--- добавили
	UpdateSong(song *model.Song) error
	DeleteSong(id uint) error
}

type songService struct {
	songRepo repository.SongRepository
}

// NewSongService создает новый сервис для песен
func NewSongService(songRepo repository.SongRepository) SongService {
	return &songService{
		songRepo: songRepo,
	}
}

// Реализация GetAllSongs
func (s *songService) GetAllSongs() ([]model.Song, error) {
	return s.songRepo.GetAll()
}

// CreateSong создает новую песню
func (s *songService) CreateSong(song *model.Song) error {
	// Бизнес-логика для создания песни
	if song.Title == "" {
		return errors.New("song title cannot be empty")
	}
	if song.Artist == "" {
		return errors.New("song artist cannot be empty")
	}
	return s.songRepo.Create(song)
}

// GetSongByID возвращает песню по ID
func (s *songService) GetSongByID(id uint) (*model.Song, error) {
	return s.songRepo.GetByID(id)
}

// GetSongsByAlbumID возвращает все песни по albumID
func (s *songService) GetSongsByAlbumID(albumID uint) ([]model.Song, error) {
	return s.songRepo.GetByAlbumID(albumID)
}

// UpdateSong обновляет информацию о песне
func (s *songService) UpdateSong(song *model.Song) error {
	// Бизнес-логика для обновления песни
	if song.Title == "" {
		return errors.New("song title cannot be empty")
	}
	if song.Artist == "" {
		return errors.New("song artist cannot be empty")
	}
	return s.songRepo.Update(song)
}

// DeleteSong удаляет песню
func (s *songService) DeleteSong(id uint) error {
	return s.songRepo.Delete(id)
}
