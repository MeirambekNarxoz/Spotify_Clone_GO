package repository

import (
	"Spotify/internal/model"
	"gorm.io/gorm"
)

// SongRepository интерфейс для работы с песнями
type SongRepository interface {
	Create(song *model.Song) error
	GetByID(id uint) (*model.Song, error)
	GetByAlbumID(albumID uint) ([]model.Song, error)
	GetAll() ([]model.Song, error)
	Update(song *model.Song) error
	Delete(id uint) error
}

// songRepository структура для работы с базой данных
type songRepository struct {
	db *gorm.DB
}

// NewSongRepository создает новый репозиторий для песен
func NewSongRepository(db *gorm.DB) SongRepository {
	return &songRepository{
		db: db,
	}
}

// GetAll возвращает все песни из базы данных
func (r *songRepository) GetAll() ([]model.Song, error) {
	var songs []model.Song
	err := r.db.Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, nil
}

// Create добавляет новую песню в базу данных
func (r *songRepository) Create(song *model.Song) error {
	return r.db.Create(song).Error
}

// GetByID находит песню по ID
func (r *songRepository) GetByID(id uint) (*model.Song, error) {
	var song model.Song
	// Используем Preload для подгрузки связанных данных (Album)
	err := r.db.Preload("Album").First(&song, id).Error
	if err != nil {
		return nil, err
	}
	return &song, nil
}

// GetByAlbumID находит все песни по AlbumID
func (r *songRepository) GetByAlbumID(albumID uint) ([]model.Song, error) {
	var songs []model.Song
	err := r.db.Where("album_id = ?", albumID).Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, nil
}

// Update обновляет данные песни
func (r *songRepository) Update(song *model.Song) error {
	return r.db.Save(song).Error
}

// Delete удаляет песню по ID
func (r *songRepository) Delete(id uint) error {
	return r.db.Delete(&model.Song{}, id).Error
}
