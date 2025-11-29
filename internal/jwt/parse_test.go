package jwt

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sso-service/internal/config"
	"sso-service/internal/domain/dto/jwt/commands"
	authErrors "sso-service/internal/domain/entities/errors"
	"testing"
	"time"
)

func TestService_Parse(t *testing.T) {
	issuer := "test-sso"
	secret := []byte("LPKCsOO6CzbXjpFUGdgZ8EtQA+oULGU+faKC60aS1Qk=")
	jwtService := service{
		secret: secret,
		config: config.SsoConfig{
			Issuer:         issuer,
			AccessTokenTTL: 15 * time.Minute,
		},
	}

	type claim struct {
		id   int64
		role string
	}

	ttl := 15 * time.Minute
	tests := []struct {
		name   string
		claims claim
		want   *claim
		ttl    time.Duration
		err    error
	}{
		{"user", claim{12, "user"}, &claim{12, "user"}, ttl, nil},
		{"expired", claim{12, "user"}, nil, time.Nanosecond, authErrors.ErrAccessTokenExpired},
		{"vip", claim{50, "vip"}, &claim{50, "vip"}, ttl, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jwtService.config.AccessTokenTTL = test.ttl
			result, err := jwtService.Generate(commands.Generate{
				UserID: test.claims.id,
				Role:   test.claims.role,
			})

			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotEmpty(t, result.Token)

			parsed, err := jwtService.Parse(commands.Parse{
				Token: result.Token,
			})

			if test.err != nil {
				assert.ErrorIs(t, err, test.err)
				assert.Nil(t, parsed)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, parsed)
			assert.Equal(t, test.want.id, parsed.UserID)
			assert.Equal(t, test.want.role, parsed.Role)
		})
	}
}
