package dbstore

type User struct {
	ID         int    `db:"id"`
	Client_ID  string `db:"client_id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Email      string `db:"email"`
	DeviceNum  string `db:"device_num"`
	DeviceType string `db:"device_type"`
	Active     bool   `db:"active"`
}
