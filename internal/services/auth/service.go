package auth

import (
	"sso-service/internal/domain/interfaces/repositories"
	"sso-service/internal/domain/interfaces/services"
)

type service struct {
	users  repositories.User
	tokens repositories.RefreshToken
}

func New(users repositories.User, tokens repositories.RefreshToken) services.Auth {
	return &service{users: users, tokens: tokens}
}
