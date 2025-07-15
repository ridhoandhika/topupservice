package domain

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	Image          string    `gorm:"type:text"`
	Title          string    `gorm:"type:varchar(255);not null"`
	Subtitle       string    `gorm:"type:varchar(255);not null"`
	Category       string    `gorm:"type:varchar(255)"`
	Provider       string    `gorm:"type:varchar(255)"`
	Sequence       int32     `gorm:"type:int"`
	Description    string    `gorm:"type:text"`
	TargetSistem   string    `gorm:"type:varchar(255)"`
	ValidatePlayer bool      `gorm:"type:boolean"`
	CodeValidation string    `gorm:"type:varchar(255)"`
	Status         bool      `gorm:"type:boolean"`
	CreatedAt      time.Time // Kolom CreatedAt
	UpdatedAt      time.Time // Kolom UpdatedAt
}

func (Game) TableName() string {
	return "topup_schema.games" // Ganti dengan nama schema yang diinginkan
}
