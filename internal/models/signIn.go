package models


type SignInReq struct{
	ClientID   string `json:"client_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	DeviceNum  string `json:"device_num"`
	DeviceType string `json:"device_type"`
	Active     bool   `json:"-"`
}

type SignInRes struct{
	AccessToken string `json:"access_toke"`
	RefreshToken string `json:"refresh_toke"`
	Active bool `json:"active"`
	ClientID string `json:"client_id"`
}