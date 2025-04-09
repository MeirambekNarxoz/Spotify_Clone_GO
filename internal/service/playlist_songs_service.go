package service

import (
	"Spotify/internal/model"
	"Spotify/internal/repository"
	"errors"
)

type PlaylistSongsService interface {
	AddSongToPlaylist(playlistID, songID uint) error                   // Добавить песню в плейлист
	RemoveSongFromPlaylist(playlistID, songID uint) error              // Удалить песню из плейлиста
	GetSongsInPlaylist(playlistID uint) ([]model.PlaylistSongs, error) // Получить все песни в плейлисте
}

type playlistSongsService struct {
	repo repository.PlaylistSongsRepository
}

func NewPlaylistSongsService(repo repository.PlaylistSongsRepository) PlaylistSongsService {
	return &playlistSongsService{repo: repo}
}

// Добавление песни в плейлист
func (s *playlistSongsService) AddSongToPlaylist(playlistID, songID uint) error {
	// Проверка, если такая связь уже существует
	existing, err := s.repo.GetByPlaylistID(playlistID)
	if err != nil {
		return err
	}

	for _, ps := range existing {
		if ps.SongID == songID {
			return errors.New("song already exists in playlist")
		}
	}

	// Создание связи
	return s.repo.Create(&model.PlaylistSongs{
		PlaylistID: playlistID,
		SongID:     songID,
	})
}

// Удаление песни из плейлиста
func (s *playlistSongsService) RemoveSongFromPlaylist(playlistID, songID uint) error {
	return s.repo.Delete(playlistID, songID)
}

// Получение всех песен в плейлисте
func (s *playlistSongsService) GetSongsInPlaylist(playlistID uint) ([]model.PlaylistSongs, error) {
	return s.repo.GetByPlaylistID(playlistID)
}
