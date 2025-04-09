package repository

import (
	"Spotify/internal/model"
	"gorm.io/gorm"
)

type PlaylistSongsRepository interface {
	Create(playlistSong *model.PlaylistSongs) error
	Delete(playlistID, songID uint) error
	GetByPlaylistID(playlistID uint) ([]model.PlaylistSongs, error)
	GetBySongID(songID uint) ([]model.PlaylistSongs, error)
}

type playlistSongsRepository struct {
	db *gorm.DB
}

func NewPlaylistSongsRepository(db *gorm.DB) PlaylistSongsRepository {
	return &playlistSongsRepository{db: db}
}

// Создание записи в таблице PlaylistSongs
func (r *playlistSongsRepository) Create(playlistSong *model.PlaylistSongs) error {
	if err := r.db.Create(playlistSong).Error; err != nil {
		return err
	}
	return nil
}

// Удаление записи по PlaylistID и SongID
func (r *playlistSongsRepository) Delete(playlistID, songID uint) error {
	if err := r.db.Where("playlist_id = ? AND song_id = ?", playlistID, songID).Delete(&model.PlaylistSongs{}).Error; err != nil {
		return err
	}
	return nil
}

// Получение песен по PlaylistID
func (r *playlistSongsRepository) GetByPlaylistID(playlistID uint) ([]model.PlaylistSongs, error) {
	var playlistSongs []model.PlaylistSongs
	if err := r.db.Where("playlist_id = ?", playlistID).Find(&playlistSongs).Error; err != nil {
		return nil, err
	}
	return playlistSongs, nil
}

// Получение песен по SongID
func (r *playlistSongsRepository) GetBySongID(songID uint) ([]model.PlaylistSongs, error) {
	var playlistSongs []model.PlaylistSongs
	if err := r.db.Where("song_id = ?", songID).Find(&playlistSongs).Error; err != nil {
		return nil, err
	}
	return playlistSongs, nil
}
