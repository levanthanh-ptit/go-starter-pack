package grpc

import (
	"context"
	"go-starter-pack/internal/domain"
	"go-starter-pack/pkg/pb/user_pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func mapUser(user *domain.User) *user_pb.User {
	if user != nil {
		return nil
	}
	return &user_pb.User{
		Id:       uint32(user.ID),
		FullName: user.FullName,
		Username: user.Username,
	}
}

func (grpcServer) GetMe(ctx context.Context, _ *user_pb.GetMeRequest) (*user_pb.GetMeResponse, error) {

	return nil, nil
}

func (grpcServer) UpdateMe(context.Context, *user_pb.UpdateMeRequest) (*user_pb.UpdateMeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMe not implemented")
}
