package repository

import (
	"context"
	"topupservice/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func User(con *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: con,
	}
}

func (u userRepository) FindByID(ctx context.Context, id uuid.UUID) (user domain.User, err error) {
	err = u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return
}
