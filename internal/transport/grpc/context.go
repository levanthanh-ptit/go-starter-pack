package grpc

import (
	"context"
	"go-starter-pack/internal/domain"
)

type contextKey int

const (
	currentUserKey contextKey = iota
)

// AppContext application context
type AppContext struct {
	context.Context
}

// AppContextFrom constructor
func AppContextFrom(ctx context.Context) *AppContext {
	return &AppContext{
		ctx,
	}
}

// WithUser bind current User to context
func (ctx *AppContext) WithUser(user *domain.User) {
	ctx.Context = context.WithValue(ctx.Context, currentUserKey, user)
}

// GetUser extract current User from context
func (ctx *AppContext) GetUser() *domain.User {
	return ctx.Value(currentUserKey).(*domain.User)
}
