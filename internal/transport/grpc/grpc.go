package grpc

import (
	"context"
	"go-starter-pack/internal/domain"
	"go-starter-pack/internal/transport"
	"go-starter-pack/pkg/pb/user_pb"
	"log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	kit "github.com/levanthanh-ptit/go-grpc-kit/server"
	"google.golang.org/grpc"
)

type grpcServer struct {
	user_pb.UnimplementedUserServiceServer
	kit.GrpcServer

	userService domain.UserService
	authService domain.AuthService
}

// NewGRPCServer constructor
func NewGRPCServer(
	name string,
	userService domain.UserService,
	authService domain.AuthService,
) transport.Server {
	base := kit.NewGrpcServer(name)
	return &grpcServer{
		GrpcServer:  *base,
		userService: userService,
		authService: authService,
	}
}

func (s *grpcServer) RegisterServer(server *grpc.Server) error {
	user_pb.RegisterUserServiceServer(server, s)
	return nil
}

func (s *grpcServer) Load(context.Context) error {
	s.WithRegister(s.RegisterServer)
	s.WithOptions(grpc_auth.UnaryServerInterceptor(s.authGuard))
	return nil
}

func (s *grpcServer) Serve(host string, port int) {
	log.Fatalln(s.ServeTCP(host, port))
}
