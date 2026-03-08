package container

import (
	"github.com/flintbits/drafenex-backend/internal/config"
	"github.com/flintbits/drafenex-backend/internal/repository"
	"github.com/flintbits/drafenex-backend/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	UserRepo      *repository.UserRepository
	OrganizerRepo *repository.OrganizerRepository

	UserService      *services.UserService
	AuthService      *services.AuthService
	OrganizerService *services.OrganizerService
}

func New(pool *pgxpool.Pool, cfg *config.Config) *Container {

	userRepo := repository.NewUserRepository(pool)
	organizerRepo := repository.NewOrganizerRepository(pool)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, cfg)
	organizerService := services.NewOrganizerService(organizerRepo)

	return &Container{
		UserRepo:      userRepo,
		OrganizerRepo: organizerRepo,

		UserService:      userService,
		AuthService:      authService,
		OrganizerService: organizerService,
	}
}
