package commands

type Base struct {
	Login    string
	Password string
	ClientID string
}

type Register struct {
	Base
}

type Login struct {
	Base
}

type Refresh struct {
	RefreshToken string
	ClientID     string
}
