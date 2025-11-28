package results

import "sso-service/internal/domain/entities"

type Generate struct {
	Token string
}

type Parse struct {
	Claims entities.JWTClaims
}
