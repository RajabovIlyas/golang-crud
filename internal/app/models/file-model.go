package models

import (
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type FileModel struct {
	ID       uuid.UUID        `json:"id"`
	FileName string           `json:"filename"`
	Path     string           `json:"path"`
	Format   database.Formats `json:"format"`
	Size     int64            `json:"size"`
	UserID   uuid.NullUUID    `json:"user_id"`
}
