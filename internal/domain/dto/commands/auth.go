package commands

type AuthBase struct {
	Login    string
	Password string
	ClientId string
}

type Register struct {
	AuthBase
}

type Login struct {
	AuthBase
}

type Refresh struct {
	RefreshToken string
	ClientId     string
}
