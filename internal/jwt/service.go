package jwt

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sso-service/internal/config"
	"sso-service/internal/domain/interfaces/services"
)

type service struct {
	secret []byte
	config config.SsoConfig
}

func New(config config.SsoConfig) (services.AccessToken, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	secret := []byte(os.Getenv("JWT_SECRET_BASE64"))
	if secret == nil {
		return nil, fmt.Errorf("failed to load JWT_SECRET_BASE64")
	}

	return &service{secret: secret, config: config}, nil
}
