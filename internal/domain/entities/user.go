package entities

var (
	RoleUser  = "user"
	RoleAdmin = "admin"
	RoleVip   = "vip"
)

type User struct {
	ID           int64
	Login        string
	PasswordHash string
	Role         string
	ClientID     string
}
