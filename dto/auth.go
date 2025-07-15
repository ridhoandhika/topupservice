package dto

type AuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResp struct {
	AccessToken string `json:"access_token"`
}
