package main

import (
	"context"
	"go-starter-pack/configs"
	"go-starter-pack/internal/domain"
	"go-starter-pack/internal/provider"
	"go-starter-pack/internal/provider/jwt"
	"go-starter-pack/internal/repository"
	"go-starter-pack/internal/repository/gormrepo"
	"go-starter-pack/internal/services"
	"go-starter-pack/internal/transport"
	"go-starter-pack/internal/transport/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/levanthanh-ptit/go-ez/ez_provider"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	config *configs.Config

	database *gorm.DB
	redisDB  *redis.Client

	bcryptProvider provider.PasswordHashProvider
	jwtProvider    provider.TokenProvider

	userRepo repository.UserRepository

	userService domain.UserService
	authService domain.AuthService

	server transport.Server
)

func initConfig() {
	config = configs.NewConfig()
}

func initProviders() {
	bcryptProvider = ez_provider.NewBcryptProvider(config.EncryptSalt, 10)
	jwtProvider = jwt.NewJWTProvider(config.JwtSigningKey, config.JwtAccessTokenDuration, config.JwtRefreshTokenDuration)
}

func initStorages(ctx context.Context) (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	database, err = gorm.Open(postgres.Open(config.DatabaseURI))
	return
}

func initRepositories() {
	userRepo = gormrepo.NewUserRepository(database)
}

func initServices() {
	userService = services.NewUserService(bcryptProvider, userRepo)
	authService = services.NewAuthService(bcryptProvider, jwtProvider, userRepo)
}

func initTransport(ctx context.Context) error {
	server = http.NewHTTPServer(userService, authService)
	if err := server.Load(ctx); err != nil {
		return err
	}
	return nil
}

func loadApp(ctx context.Context) error {
	initConfig()
	initProviders()
	if err := initStorages(ctx); err != nil {
		return err
	}
	initRepositories()
	initServices()
	if err := initTransport(ctx); err != nil {
		return err
	}
	return nil
}

func start() {
	server.Serve(config.Host, config.Port)
}

func main() {
	ctx := context.Background()
	if err := loadApp(ctx); err != nil {
		panic(err)
	}
	start()
}
