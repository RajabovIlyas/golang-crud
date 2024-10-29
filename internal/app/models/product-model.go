package models

import (
	"github.com/google/uuid"
	"time"
)

type Products struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Name         string    `json:"name" gorm:"unique;not null"`
	Year         int       `json:"year"`
	Price        float64   `json:"price"`
	CPUModel     string    `json:"cpu_model"`
	HardDiskSize string    `json:"hard_disk_size"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}
