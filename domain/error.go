package domain

import "errors"

var ErrAuthFailed = errors.New("error authentication failed")
var UserAlreadyExist = errors.New("user already exists")

type BaseResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
