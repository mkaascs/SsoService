package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"sso-service/internal/config"
	"sso-service/internal/domain/interfaces/repositories"
	"sso-service/internal/domain/interfaces/services"
	"sso-service/internal/lib/rand"
)

type service struct {
	users      repositories.User
	tokens     repositories.RefreshToken
	jwt        services.JWT
	log        *slog.Logger
	config     config.SsoConfig
	hmacSecret []byte
}

func (s *service) generateRefreshToken() string {
	token, _ := rand.GenerateBase64(32)
	return token
}

func (s *service) hashRefreshToken(token string) string {
	h := hmac.New(sha256.New, s.hmacSecret)
	h.Write([]byte(token))
	return hex.EncodeToString(h.Sum(nil))
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
