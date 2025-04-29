package model

import "gorm.io/gorm"

// TODO json serialization
type Song struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title"`
	Artist   string `gorm:"not null" json:"artist"`            // Просто текст, как вы хотели
	Audio    string `gorm:"not null" json:"audio"`             // Путь к аудиофайлу или URL
	CoverImg string `gorm:"column:cover_img" json:"cover_img"` // Обложка песни
	AlbumID  uint   `gorm:"not null" json:"album_id"`          // Альбом, к которому относится песня
	Album    Album  `gorm:"foreignKey:AlbumID" `               // Связь с альбомом
}
