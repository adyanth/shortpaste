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
	ID   string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	Link string `json:"link" validate:"required,url"`
	gorm.Model
}

type File struct {
	ID   string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	Name string `json:"name"`
	MIME string
	gorm.Model
}

type Text struct {
	ID   string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	Type string `validate:"omitempty,oneof=txt md"`
	Text string `gorm:"-" json:"text"`
	gorm.Model
}
