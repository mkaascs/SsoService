package entities

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type JWTClaims struct {
	UserID int64
}
