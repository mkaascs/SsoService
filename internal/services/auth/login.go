package auth

import (
	"context"
	"sso-service/internal/domain/dto/commands"
	"sso-service/internal/domain/dto/results"
)

func (s *service) Login(ctx context.Context, command commands.Login) (*results.Login, error) {
	return nil, nil
}
