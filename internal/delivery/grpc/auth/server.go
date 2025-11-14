package auth

import (
	ssov1 "github.com/mkaascs/SsoProto/gen/go/sso"
	"google.golang.org/grpc"
)

type server struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &server{})
}
