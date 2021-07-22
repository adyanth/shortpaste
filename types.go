package shortpaste

import (
	"embed"

	"gorm.io/gorm"
)

// Embed html templates with code
//go:embed templates/*
var templateFS embed.FS

// App struct containing the bind address, storage path and the db connector.
type App struct {
	bind            string
	db              *gorm.DB
	storagePath     string
	link307Redirect bool
	username        string
	password        string
	noAuth          bool
}

// Link struct for saving the Redirect Links /l/.
type Link struct {
	ID       string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	Link     string `json:"link" validate:"required,url"`
	HitCount int64  `json:"hitcount" validate:"isdefault"`
	gorm.Model
}

// File struct for saving the file uploads /f/.
type File struct {
	ID            string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	Name          string `json:"name"`
	MIME          string `validate:"isdefault"`
	HitCount      int64  `json:"hitcount" validate:"isdefault"`
	DownloadCount int64  `json:"downloadcount" validate:"isdefault"`
	gorm.Model
}

// Text struct for saving the text pastes /t/.
type Text struct {
	ID          string `gorm:"primaryKey" json:"id" validate:"required,min=3,max=32,alphanumunicode"`
	Type        string `validate:"omitempty,oneof=txt md" json:"type"`
	Text        string `gorm:"-" json:"text,omitempty"`
	NoHighlight bool   `json:"nohighlight"`
	HitCount    int64  `json:"hitcount" validate:"isdefault"`
	gorm.Model
}
