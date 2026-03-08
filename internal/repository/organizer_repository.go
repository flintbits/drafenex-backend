package repository

import (
	"context"

	"github.com/flintbits/drafenex-backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganizerRepository struct {
	pool *pgxpool.Pool
}

func NewOrganizerRepository(pool *pgxpool.Pool) *OrganizerRepository {
	return &OrganizerRepository{pool: pool}
}

func (r *OrganizerRepository) CreateOrganizer(ctx context.Context, organizer *models.Organizer) (*models.Organizer, error) {
	query := `
		INSERT INTO organizers (
			user_id,
			full_name,
			phone_number,
			company_name,
			avatar_url,
			website,
			bio,
			address_line1,
			address_line2,
			city,
			state,
			country,
			postal_code
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
		RETURNING
			id,
			user_id,
			full_name,
			phone_number,
			company_name,
			avatar_url,
			website,
			bio,
			address_line1,
			address_line2,
			city,
			state,
			country,
			postal_code,
			status,
			is_verified,
			created_at,
			updated_at
	`

	created := models.Organizer{}

	err := r.pool.QueryRow(ctx, query,
		organizer.UserID,
		organizer.FullName,
		organizer.PhoneNumber,
		organizer.CompanyName,
		organizer.AvatarUrl,
		organizer.Website,
		organizer.Bio,
		organizer.AddressLine1,
		organizer.AddressLine2,
		organizer.City,
		organizer.State,
		organizer.Country,
		organizer.PostalCode,
	).Scan(
		&created.ID,
		&created.UserID,
		&created.FullName,
		&created.PhoneNumber,
		&created.CompanyName,
		&created.AvatarUrl,
		&created.Website,
		&created.Bio,
		&created.AddressLine1,
		&created.AddressLine2,
		&created.City,
		&created.State,
		&created.Country,
		&created.PostalCode,
		&created.Status,
		&created.IsVerified,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &created, nil
}
