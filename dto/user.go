package dto

import "github.com/google/uuid"

type UserData struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"fullname"`
	Email    string    `json:"email"`
}

type UserRegisterReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterResp struct {
	Message string `json:"message"`
}
