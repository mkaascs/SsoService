package auth

import "sso-service/internal/domain/interfaces/repositories"

type Service struct {
	users  repositories.User
	tokens repositories.RefreshToken
}

func New(users repositories.User, tokens repositories.RefreshToken) *Service {
	return &Service{users: users, tokens: tokens}
}
