package dto

import "github.com/google/uuid"

type UserData struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"fullname"`
	Email    string    `json:"email"`
}
