package service

import (
	"context"
	"topupservice/domain"
	"topupservice/dto"

	"github.com/google/uuid"
)

type userService struct {
	userRepository domain.UserRepository
}

func User(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) GetUser(ctx context.Context, userID string) (interface{}, error) {
	parsedUserID, _ := uuid.Parse(userID)
	user, err := u.userRepository.FindByID(ctx, parsedUserID)
	if err != nil {
		return nil, err
	}

	userResp := dto.UserData{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return userResp, nil
}
