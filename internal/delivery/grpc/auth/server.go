package auth

import (
	ssov1 "github.com/mkaascs/SsoProto/gen/go/sso"
	"google.golang.org/grpc"
	"sso-service/internal/domain/services"
)

type server struct {
	ssov1.UnimplementedAuthServer
	auth services.Auth
}

func Register(gRPC *grpc.Server, auth services.Auth) {
	ssov1.RegisterAuthServer(gRPC, &server{auth: auth})
}
