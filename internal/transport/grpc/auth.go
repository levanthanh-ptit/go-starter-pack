package grpc

import (
	"context"
	"go-starter-pack/internal/domain"
	"go-starter-pack/pkg/pb/user_pb"
	"strings"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

var (
	authWhiteList = map[string]bool{
		"/user_pb.UserService/Register": true,
		"/user_pb.UserService/Login":    true,
	}
)

func (s *grpcServer) authGuard(ctx context.Context) (context.Context, error) {
	if method, _ := grpc.Method(ctx); authWhiteList[method] {
		return ctx, nil
	}
	appCtx := AppContextFrom(ctx)
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, domain.ErrUnauthorized
	}
	result, err := s.authService.VerifyToken(ctx, token)
	if err != nil {
		return nil, err
	}
	appCtx.WithUser(result.Payload)
	return appCtx, nil
}

func (s *grpcServer) Register(ctx context.Context, u *user_pb.RegisterRequest) (*user_pb.RegisterResponse, error) {
	if _, err := s.userService.CreateUser(ctx, &domain.User{
		FullName: strings.Trim(u.FullName, " "),
		Username: strings.Trim(u.Username, " "),
		Password: u.Password,
	}); err != nil {
		return nil, err
	}
	return new(user_pb.RegisterResponse), nil
}
func (s *grpcServer) Login(ctx context.Context, u *user_pb.LoginRequest) (*user_pb.LoginResponse, error) {
	result, err := s.authService.Login(ctx, &domain.LoginCredential{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}

	return &user_pb.LoginResponse{
		Profile: mapUser(result.Profile),
		Token:   result.Token,
	}, nil
}
