package models

type SignInReq struct {
	ClientID    string `json:"client_id" example:"12345"`
	Email       string `json:"email" example:"@hamkorbank.uz"`
	AccessToken string `json:"access_token" example:"valid token"`
}
