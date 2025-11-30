package commands

type Add struct {
	Login        string
	Role         string
	PasswordHash string
	ClientID     string
}

type GetByLogin struct {
	Login    string
	ClientID string
}
