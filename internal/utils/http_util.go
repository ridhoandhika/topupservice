package util

import (
	"errors"
	"topupservice/domain"
)

func GetHttpStatus(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	case errors.Is(err, domain.UserAlreadyExist):
		return 409
	default:
		return 500
	}
}
