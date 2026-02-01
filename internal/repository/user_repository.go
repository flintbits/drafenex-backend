package repository

import (
	"context"

	"github.com/flintbits/drafenex-backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	query := `
		INSERT INTO users (email, password_hash,role)
		VALUES ($1,$2,$3)
		RETURNING id,email,role,is_onboarded,created_at,updated_at
	`

	created := models.User{}

	err := r.pool.QueryRow(ctx, query, user.Email, user.PasswordHash, user.Role).Scan(
		&created.ID,
		&created.Email,
		&created.Role,
		&created.IsOnboarded,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id,email,password_hash,role,is_onboarded,created_at,updated_at
		FROM users
		WHERE email = $1
	`
	fetchedUser := models.User{}

	err := r.pool.QueryRow(ctx, query, email).Scan(
		&fetchedUser.ID,
		&fetchedUser.Email,
		&fetchedUser.PasswordHash,
		&fetchedUser.Role,
		&fetchedUser.IsOnboarded,
		&fetchedUser.CreatedAt,
		&fetchedUser.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &fetchedUser, nil
}
