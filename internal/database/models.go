// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Formats string

const (
	FormatsVideo Formats = "video"
	FormatsPhoto Formats = "photo"
	FormatsMusic Formats = "music"
	FormatsOther Formats = "other"
)

func (e *Formats) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Formats(s)
	case string:
		*e = Formats(s)
	default:
		return fmt.Errorf("unsupported scan type for Formats: %T", src)
	}
	return nil
}

type NullFormats struct {
	Formats Formats
	Valid   bool // Valid is true if Formats is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFormats) Scan(value interface{}) error {
	if value == nil {
		ns.Formats, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Formats.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFormats) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Formats), nil
}

type File struct {
	ID        uuid.UUID
	FileName  string
	Path      string
	Format    Formats
	Size      int64
	UserID    uuid.NullUUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Token struct {
	ID             uuid.UUID
	AccessTokenKey uuid.UUID
	UserID         uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type User struct {
	ID        uuid.UUID
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Password  string
}
