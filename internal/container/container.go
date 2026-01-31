package container

import (
	"github.com/flintbits/drafenex-backend/internal/repository"
	"github.com/flintbits/drafenex-backend/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	UserRepo *repository.UserRepository

	UserService *services.UserService
}

func New(pool *pgxpool.Pool) *Container {
	userRepo := repository.NewUserRepository(pool)

	userService := services.NewUserService(userRepo)

	return &Container{
		UserRepo:    userRepo,
		UserService: userService,
	}
}
