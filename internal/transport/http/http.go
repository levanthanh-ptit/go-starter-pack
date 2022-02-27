package http

import (
	"context"
	"fmt"
	"go-starter-pack/internal/domain"
	"go-starter-pack/internal/transport"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// httpServer http transport
type httpServer struct {
	echo        *echo.Echo
	userService domain.UserService
	authService domain.AuthService
}

// NewHTTPServer constructor
func NewHTTPServer(
	userService domain.UserService,
	authService domain.AuthService,
) transport.Server {
	return &httpServer{
		echo.New(),
		userService,
		authService,
	}
}

func (s *httpServer) Load(_ context.Context) error {
	s.echo.Use(middleware.Logger())
	s.echo.Use(httpErrorResolver)
	s.echo.POST("/auth/register", s.Register)
	s.echo.POST("/auth/login", s.Login)
	s.echo.GET("/users/me", s.GetMe, s.authGuard)
	s.echo.PATCH("/users/me", s.UpdateMe, s.authGuard)
	return nil
}

func (s *httpServer) Serve(host string, port int) {
	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf("%s:%v", host, port)))
}
