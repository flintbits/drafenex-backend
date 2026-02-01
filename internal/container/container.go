package container

import (
	"github.com/flintbits/drafenex-backend/internal/config"
	"github.com/flintbits/drafenex-backend/internal/repository"
	"github.com/flintbits/drafenex-backend/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	UserRepo *repository.UserRepository

	UserService *services.UserService
	AuthService *services.AuthService
}

func New(pool *pgxpool.Pool, cfg *config.Config) *Container {

	userRepo := repository.NewUserRepository(pool)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, cfg)

	return &Container{
		UserRepo:    userRepo,
		UserService: userService,
		AuthService: authService,
	}
}
