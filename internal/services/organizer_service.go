package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/flintbits/drafenex-backend/internal/dto"
	"github.com/flintbits/drafenex-backend/internal/models"
	"github.com/flintbits/drafenex-backend/internal/repository"
)

type OrganizerService struct {
	organizerRepo *repository.OrganizerRepository
}

func NewOrganizerService(organizerRepo *repository.OrganizerRepository) *OrganizerService {
	return &OrganizerService{organizerRepo: organizerRepo}
}

func (s *OrganizerService) CreateOrganizer(ctx context.Context, UserID int64, input *dto.CreateOrganizerInput) (*models.Organizer, error) {

	organizer := &models.Organizer{
		UserID:      UserID,
		FullName:    input.FullName,
		PhoneNumber: input.PhoneNumber,

		CompanyName: input.CompanyName,
		AvatarUrl:   input.AvatarUrl,
		Website:     input.Website,
		Bio:         input.Bio,

		AddressLine1: input.AddressLine1,
		AddressLine2: input.AddressLine2,
		City:         input.City,
		State:        input.State,
		Country:      input.Country,
		PostalCode:   input.PostalCode,
		Status:       models.OrganizerPending,
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	createdOrganizer, err := s.organizerRepo.CreateOrganizer(ctx, organizer)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			return nil, errors.New("Organizer already registered")
		}

		return nil, err
	}

	return createdOrganizer, nil
}
