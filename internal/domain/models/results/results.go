package results

import "sso-service/internal/domain/models"

type Register struct {
	UserId int64
}

type Login struct {
	Tokens *models.TokenPair
}

type Refresh struct {
	Tokens *models.TokenPair
}
