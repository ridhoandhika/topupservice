package domain

import (
	"context"
	"time"
	"topupservice/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`       // UUID sebagai primary key
	Username  string    `gorm:"type:varchar(100);not null"` // Kolom Username yang unik dan tidak boleh kosong
	Password  string    `gorm:"type:varchar(255);not null"` // Kolom Password yang tidak boleh kosong
	Email     string    `gorm:"type:varchar(255);not null"` // Kolom Password yang tidak boleh kosong
	CreatedAt time.Time // Kolom CreatedAt
	UpdatedAt time.Time // Kolom UpdatedAt
}

func (User) TableName() string {
	return "topup_schema.users" // Ganti dengan nama schema yang diinginkan
}

type UserRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	InsertUser(ctx context.Context, req dto.UserRegisterReq) (bool, error)
}

type UserService interface {
	GetUser(ctx context.Context, userID string) (interface{}, error)
}
