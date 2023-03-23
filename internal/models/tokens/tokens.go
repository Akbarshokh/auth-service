package tokens

type Token struct {
	AccessToken string `json:"access_toke"`
	RefreshToken string `json:"refresh_toke"`
	Active bool `json:"active"`
	ClientID string `json:"client_id"`
}
