package dto

import "github.com/google/uuid"

type UserData struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"fullname"`
	Level    string    `json:"level"`
}

type UserRegisterReq struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"123456"`
	Level    string `json:"level" enums:"Superadmin,CS,Admin,Member,Silver,Gold" example:"Member"`
	Whatsapp string `json:"whatsapp" example:"+628123456789"`
}

type UserRegisterResp struct {
	Message string `json:"message"`
}
