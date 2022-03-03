package grpc

import (
	"context"
	"fmt"
	"go-starter-pack/internal/transport"
	"go-starter-pack/pkg/pb/user_pb"
	"log"

	kit "github.com/levanthanh-ptit/go-grpc-kit/server"
	"google.golang.org/grpc"
)

type grpcServer struct {
	user_pb.UnimplementedUserServiceServer
	kit.GrpcServer
}

// NewGRPCServer constructor
func NewGRPCServer(name string) transport.Server {
	base := kit.NewGrpcServer(name)
	return &grpcServer{
		GrpcServer: *base,
	}
}

func (s *grpcServer) RegisterServer(server *grpc.Server) error {
	user_pb.RegisterUserServiceServer(server, s)
	return nil
}

func (s *grpcServer) Load(context.Context) error {
	s.WithGrpcRegister(s.RegisterServer)
	return nil
}

func (s *grpcServer) Serve(host string, port int) {
	log.Fatalln(s.WithHost(host).WithPort(fmt.Sprint(port)).ServeTCP())
}
