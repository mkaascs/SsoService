package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sso-service/internal/config"
	"sso-service/internal/domain/dto/commands"
	"testing"
	"time"
)

func TestService_Generate(t *testing.T) {
	issuer := "test-sso"
	secret := []byte("93RmM4BWwZ4WhhF1n7+Pl1w1nPKlVbNMPWcj6jFrV7Y=")
	jwtService := &service{
		secret: secret,
		config: config.SsoConfig{
			AccessTokenTTL: 15 * time.Minute,
			Issuer:         issuer,
		},
	}

	tests := []struct {
		id   int64
		role string
	}{
		{54, "user"},
		{20, "vip"},
		{2, "admin"},
	}

	for _, test := range tests {
		t.Run(test.role, func(t *testing.T) {
			result, err := jwtService.Generate(commands.Generate{
				UserID: test.id,
				Role:   test.role,
			})

			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotEmpty(t, result.Token)

			claims := &Claims{}
			token, err := jwt.ParseWithClaims(result.Token, claims, func(token *jwt.Token) (any, error) {
				return secret, nil
			})

			require.NoError(t, err)
			require.NotNil(t, token)
			require.True(t, token.Valid)

			assert.Equal(t, test.id, claims.UserID)
			assert.Equal(t, test.role, claims.Role)
			assert.Equal(t, issuer, claims.Issuer)
		})
	}
}
