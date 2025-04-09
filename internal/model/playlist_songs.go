package model

import "gorm.io/gorm"

type PlaylistSongs struct {
	gorm.Model
	PlaylistID uint `gorm:"not null" json:"playlist_id"`
	SongID     uint `gorm:"not null" json:"song_id"`

	Playlist Playlist `gorm:"foreignKey:PlaylistID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Song     Song     `gorm:"foreignKey:SongID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
