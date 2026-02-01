package services

import (
	"context"
	"errors"
	"time"

	"github.com/flintbits/drafenex-backend/internal/auth"
	"github.com/flintbits/drafenex-backend/internal/config"
	"github.com/flintbits/drafenex-backend/internal/models"
	"github.com/flintbits/drafenex-backend/internal/repository"
)

type AuthService struct {
	userRepo *repository.UserRepository
	cfg      *config.Config
}

func NewAuthService(userRepo *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (string, *models.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", nil, errors.New("invalid credentails 1")
	}

	if err := auth.ComparePassword(user.PasswordHash, password); err != nil {
		return "", nil, errors.New("invalid credentials 2")
	}

	token, err := auth.GenerateAccessToken(
		user.ID,
		user.Email,
		user.Role,
		s.cfg.JWTSecret,
		24*time.Hour,
	)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}
