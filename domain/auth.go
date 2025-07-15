package domain

import (
	"context"
	"topupservice/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthReq) (*dto.AuthResp, error)
	Refresh(ctx context.Context, token string) (*dto.AuthResp, error)
	Register(ctx context.Context, req dto.UserRegisterReq) (bool, error)
}
