package auth

import (
	"context"
	"sso-service/internal/domain/dto/commands"
	"sso-service/internal/domain/dto/results"
)

func (s *service) Refresh(ctx context.Context, command commands.Refresh) (*results.Refresh, error) {
	return nil, nil
}
