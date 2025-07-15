package repository

import (
	"context"
	"topupservice/domain"
	"topupservice/dto"

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

func (u userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	err = u.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return
}

func (u userRepository) InsertUser(ctx context.Context, req dto.UserRegisterReq) (bool, error) {
	user := domain.User{
		ID:       uuid.New(),
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}
