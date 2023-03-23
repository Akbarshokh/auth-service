package user

type User struct {
	ClientID   string `json:"client_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	DeviceNum  string `json:"device_num"`
	DeviceType string `json:"device_type"`
	Active     bool   `json:"-"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
