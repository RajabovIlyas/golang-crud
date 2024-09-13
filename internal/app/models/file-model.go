package models

import (
	"github.com/google/uuid"
	"time"
)

type Formats string

const (
	FormatsVideo Formats = "video"
	FormatsPhoto Formats = "photo"
	FormatsMusic Formats = "music"
	FormatsOther Formats = "other"
)

type CreateFile struct {
	FileName string        `json:"file_name" gorm:"not null"`
	Path     string        `json:"path" gorm:"not null"`
	Format   Formats       `json:"format" gorm:"not null"`
	Size     int64         `json:"size" gorm:"not null"`
	UserID   uuid.NullUUID `json:"user_id" `
}

type ResponseFile struct {
	Url string `json:"url"`
}

type Files struct {
	ID        uuid.UUID     `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	FileName  string        `json:"file_name" gorm:"not null"`
	Path      string        `json:"path" gorm:"not null"`
	Format    Formats       `json:"format" gorm:"not null"`
	Size      int64         `json:"size" gorm:"not null"`
	UserID    uuid.NullUUID `json:"user_id" `
	CreatedAt time.Time     `json:"-" gorm:"not null;default:now()"`
	UpdatedAt time.Time     `json:"-" gorm:"not null;default:now()"`
}
