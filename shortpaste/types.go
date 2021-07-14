package shortpaste

import (
	"gorm.io/gorm"
)

type App struct {
	bind        string
	db          *gorm.DB
	storagePath string
}

type Link struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Link string `json:"link"`
	gorm.Model
}

type File struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Path string `json:"path"`
	gorm.Model
}

type Text struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Path string
	Text string `gorm:"-" json:"text"`
	gorm.Model
}
