package auth

import (
	"context"
	"errors"
	ssov1 "github.com/mkaascs/SsoProto/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sso-service/internal/domain/models/commands"
	authErrors "sso-service/internal/domain/models/errors"
)

func (s *server) Login(ctx context.Context, request *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	// TODO: validate

	result, err := s.auth.Login(ctx, commands.Login{
		AuthBase: commands.AuthBase{
			Login:    request.Login,
			Password: request.Password,
			ClientId: request.ClientId,
		},
	})

	if err != nil {
		if errors.Is(err, authErrors.ErrUserNotFound) || errors.Is(err, authErrors.ErrInvalidPassword) {
			return nil, status.Error(codes.Unauthenticated, "invalid login or password")
		}

		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &ssov1.LoginResponse{
		AccessToken:  result.Tokens.AccessToken,
		RefreshToken: result.Tokens.RefreshToken,
	}, nil
}
