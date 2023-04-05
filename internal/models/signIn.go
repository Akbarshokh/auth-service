package models

type SignInReq struct {
	ClientID    string `json:"client_id" example:"12345"`
	Email       string `json:"email" example:"@hamkorbank.uz"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token" example:"valid token"`
}

type SignInRes struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	AccessToken string `json:"access_token"`
}