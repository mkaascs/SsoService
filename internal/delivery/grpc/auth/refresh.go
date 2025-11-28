package auth

import (
	"context"
	"errors"
	ssov1 "github.com/mkaascs/SsoProto/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sso-service/internal/domain/dto/commands"
	authErrors "sso-service/internal/domain/entities/errors"
)

func (s *server) Refresh(ctx context.Context, request *ssov1.RefreshRequest) (*ssov1.RefreshResponse, error) {
	// TODO: validate

	result, err := s.auth.Refresh(ctx, commands.Refresh{
		RefreshToken: request.RefreshToken,
		ClientId:     request.ClientId,
	})

	if err != nil {
		if errors.Is(err, authErrors.ErrInvalidRefreshToken) {
			return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
		}

		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &ssov1.RefreshResponse{
		AccessToken:  result.Tokens.AccessToken,
		RefreshToken: result.Tokens.RefreshToken,
	}, nil
}
