package grpc

import (
	"context"
	"go-starter-pack/pkg/pb/user_pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpcServer) CreateUser(context.Context, *user_pb.CreateUserRequest) (*user_pb.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (grpcServer) GetUser(context.Context, *user_pb.GetUserRequest) (*user_pb.GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func (grpcServer) ListUser(context.Context, *user_pb.ListUserRequest) (*user_pb.ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}

func (grpcServer) Login(context.Context, *user_pb.LoginRequest) (*user_pb.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func (grpcServer) VerifyAuthToken(context.Context, *user_pb.VerifyAuthTokenRequest) (*user_pb.VerifyAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAuthToken not implemented")
}
