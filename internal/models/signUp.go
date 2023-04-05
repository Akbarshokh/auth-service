package models

type SignUpReq struct {
	ClientID   string `json:"client_id" example:"12345"`
	FirstName  string `json:"first_name" example:"Ism"`
	LastName   string `json:"last_name" example:"Familiya" `
	Email      string `json:"email" example:"@hamkorbank.uz"`
	Password   string `json:"password" example:"Password@"`
	DeviceNum  string `json:"device_num" example:"172.25.102.25 / 423265"`
	DeviceType string `json:"device_type" example:"web / mobile"`
	Active     bool   `json:"-"`
}

type SignUpRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Active       bool   `json:"active"`
	ClientID     string `json:"client_id"`
}
