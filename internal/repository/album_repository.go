package repository

import (
	"Spotify/internal/model"
	"gorm.io/gorm"
)

// AlbumRepository интерфейс для работы с альбомами
type AlbumRepository interface {
	Create(album *model.Album) error
	GetByID(id uint) (*model.Album, error)
	GetAll() ([]model.Album, error)
	Update(album *model.Album) error
	Delete(id uint) error
}

// albumRepository структура для работы с базой данных
type albumRepository struct {
	db *gorm.DB
}

// NewAlbumRepository создает новый репозиторий для альбомов
func NewAlbumRepository(db *gorm.DB) AlbumRepository {
	return &albumRepository{
		db: db,
	}
}

// Create добавляет новый альбом в базу данных
func (r *albumRepository) Create(album *model.Album) error {
	return r.db.Create(album).Error
}

// GetByID находит альбом по ID
func (r *albumRepository) GetByID(id uint) (*model.Album, error) {
	var album model.Album
	err := r.db.First(&album, id).Error
	if err != nil {
		return nil, err
	}
	return &album, nil
}

// GetAll находит все альбомы
func (r *albumRepository) GetAll() ([]model.Album, error) {
	var albums []model.Album
	err := r.db.Find(&albums).Error
	if err != nil {
		return nil, err
	}
	return albums, nil
}

// Update обновляет данные альбома
func (r *albumRepository) Update(album *model.Album) error {
	return r.db.Save(album).Error
}

// Delete удаляет альбом по ID
func (r *albumRepository) Delete(id uint) error {
	return r.db.Delete(&model.Album{}, id).Error
}
