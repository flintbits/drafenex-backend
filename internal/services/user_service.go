package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/flintbits/drafenex-backend/internal/models"
	"github.com/flintbits/drafenex-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, email, password, role string) (*models.User, error) {
	if len(password) < 6 {
		return nil, errors.New("Password must be atleast 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         role,
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	createdUser, err := s.userRepo.CreateUser(ctx, user)

	if err != nil {
		if strings.Contains(err.Error(), "duplucate") || strings.Contains(err.Error(), "unique") {
			return nil, errors.New("Email already registered")
		}

		return nil, err
	}

	return createdUser, nil
}
