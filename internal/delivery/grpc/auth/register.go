package auth

import (
	"context"
	"errors"
	ssov1 "github.com/mkaascs/SsoProto/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sso-service/internal/domain/dto/auth/commands"
	authErrors "sso-service/internal/domain/entities/errors"
)

func (s *server) Register(ctx context.Context, request *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	// TODO: validate

	result, err := s.auth.Register(ctx, commands.Register{
		Base: commands.Base{
			Login:    request.Login,
			Password: request.Password,
			ClientID: request.ClientId,
		},
	})

	if err != nil {
		if errors.Is(err, authErrors.ErrUserAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "user with this login already exists")
		}

		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &ssov1.RegisterResponse{
		UserId: result.UserID,
	}, nil
}
