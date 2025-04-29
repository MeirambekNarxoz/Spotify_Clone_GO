package model

import "gorm.io/gorm"

// TODO json serialization
type Playlist struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	CoverImg string `gorm:"column:cover_img" json:"cover_img"`
	UserID   uint   `gorm:"not null" json:"userID"`                          // Владелец плейлиста (только ID пользователя)
	Songs    []Song `gorm:"many2many:playlist_songs;"json:"songs,omitempty"` // Связь с песнями через промежуточную таблицу
}
