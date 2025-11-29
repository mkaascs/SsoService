package commands

import "time"

type Add struct {
	UserID           int64
	RefreshTokenHash string
	ClientID         string
	ExpiresAt        time.Time
}

type Update struct {
	RefreshTokenHash    string
	NewRefreshTokenHash string
	ClientID            string
}
