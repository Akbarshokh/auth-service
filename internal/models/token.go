package models

type CheckTokenReq struct {
	AccessToken string `json:"access_token" example:"valid token" `
}

type CheckTokenRes struct {
	Active bool `json:"active"`
}

type GetTokenReq struct {
	RefreshToken string `json:"refresh_token" example:"valid token"`
}

type GetTokenRes struct {
	SignUpRes
}
