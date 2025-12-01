package entities

var (
	WebAppClientID = "web-app"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}
