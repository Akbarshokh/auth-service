package models

type SignInReq struct {
	ClientID    string `json:"client_id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
