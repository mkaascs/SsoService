package commands

import "time"

type Add struct {
	UserID           int64
	RefreshTokenHash string
	ClientID         string
	ExpiresAt        time.Time
}

type UpdateByToken struct {
	RefreshTokenHash    string
	NewRefreshTokenHash string
	ClientID            string
}

type UpdateByUserID struct {
	UserID              int64
	NewRefreshTokenHash string
}
