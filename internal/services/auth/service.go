package auth

import (
	"log/slog"
	"sso-service/internal/config"
	"sso-service/internal/domain/interfaces/repositories"
	"sso-service/internal/domain/interfaces/services"
)

type service struct {
	users      repositories.User
	tokens     repositories.RefreshToken
	jwt        services.JWT
	log        *slog.Logger
	config     config.SsoConfig
	hmacSecret []byte
}

func New(users repositories.User, tokens repositories.RefreshToken, jwt services.JWT, log *slog.Logger, config config.SsoConfig, hmacSecret []byte) services.Auth {
	return &service{
		users:      users,
		tokens:     tokens,
		jwt:        jwt,
		log:        log,
		config:     config,
		hmacSecret: hmacSecret,
	}
}
