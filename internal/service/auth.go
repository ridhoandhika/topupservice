package service

import (
	"context"
	"errors"
	"fmt"
	"topupservice/domain"
	"topupservice/dto"
	util "topupservice/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepository domain.UserRepository
}

func Auth(userRepository domain.UserRepository) domain.AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (u authService) Login(ctx context.Context, req dto.AuthReq) (*dto.AuthResp, error) {

	// get user by username
	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, domain.ErrAuthFailed
	}

	// Bandingkan password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, domain.ErrAuthFailed
	}

	// generate jwt
	token, err := util.GenerateTokenJWT(req.Username, user.ID.String())
	fmt.Println("error login token", token)
	if err != nil {
		return nil, domain.ErrAuthFailed
	}

	return &dto.AuthResp{
		AccessToken: token,
	}, nil

}

func (u authService) Refresh(ctx context.Context, token string) (*dto.AuthResp, error) {
	// Memverifikasi token menggunakan VerifyToken
	tokenResp, err := util.VerifyToken(token)

	if err != nil {
		// Jika token tidak valid, mengembalikan response error
		return nil, errors.New("invalid")
	}

	claims, ok := tokenResp.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid")
	}

	// Mengambil nilai 'username' dari klaim
	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("invalid")
	}

	// Mengambil nilai 'username' dari klaim
	userId, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("invalid")
	}

	newToken, err := util.GenerateTokenJWT(username, userId)
	if err != nil {
		return nil, errors.New("invalid")
	}

	return &dto.AuthResp{
		AccessToken: newToken,
	}, nil
}

func (u authService) Register(ctx context.Context, req dto.UserRegisterReq) (bool, error) {

	existingUser, err := u.userRepository.FindByUsername(ctx, req.Username)

	if err == nil && existingUser.Username != "" {
		return false, domain.UserAlreadyExist
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, errors.New("failed to hash password")
	}
	req.Password = string(hashedPassword)

	_, err = u.userRepository.InsertUser(ctx, req)
	if err != nil {
		return false, errors.New("failed to create user")
	}

	return true, nil
}
