package model

import "gorm.io/gorm"

// TODO json serialization
type Album struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	CoverImg string `gorm:"column:cover_img" json:"cover_img"`         // Обложка альбома
	Songs    []Song `gorm:"foreignKey:AlbumID" json:"songs,omitempty"` // Связь с песнями
}
