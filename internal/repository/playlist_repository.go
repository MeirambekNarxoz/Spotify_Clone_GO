package repository

import (
	"Spotify/internal/model"
	"gorm.io/gorm"
)

// PlaylistRepository интерфейс для работы с плейлистами
type PlaylistRepository interface {
	Create(playlist *model.Playlist) error
	GetByID(id uint) (*model.Playlist, error)
	GetByUserID(userID uint) ([]model.Playlist, error)
	GetAll() ([]model.Playlist, error)
	Update(playlist *model.Playlist) error
	Delete(id uint) error
}

// playlistRepository структура для работы с базой данных
type playlistRepository struct {
	db *gorm.DB
}

// NewPlaylistRepository создает новый репозиторий для плейлистов
func NewPlaylistRepository(db *gorm.DB) PlaylistRepository {
	return &playlistRepository{
		db: db,
	}
}

// GetAll возвращает все плейлисты
func (r *playlistRepository) GetAll() ([]model.Playlist, error) {
	var playlists []model.Playlist
	err := r.db.Find(&playlists).Error
	if err != nil {
		return nil, err
	}
	return playlists, nil
}

// Create добавляет новый плейлист в базу данных
func (r *playlistRepository) Create(playlist *model.Playlist) error {
	return r.db.Create(playlist).Error
}

// GetByID находит плейлист по ID
func (r *playlistRepository) GetByID(id uint) (*model.Playlist, error) {
	var playlist model.Playlist
	err := r.db.First(&playlist, id).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

// GetByUserID находит все плейлисты для пользователя
func (r *playlistRepository) GetByUserID(userID uint) ([]model.Playlist, error) {
	var playlists []model.Playlist
	err := r.db.Where("user_id = ?", userID).Find(&playlists).Error
	if err != nil {
		return nil, err
	}
	return playlists, nil
}

// Update обновляет данные плейлиста
func (r *playlistRepository) Update(playlist *model.Playlist) error {
	return r.db.Save(playlist).Error
}

// Delete удаляет плейлист по ID
func (r *playlistRepository) Delete(id uint) error {
	return r.db.Delete(&model.Playlist{}, id).Error
}
