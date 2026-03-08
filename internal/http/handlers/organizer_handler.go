package handlers

import (
	"net/http"

	"github.com/flintbits/drafenex-backend/internal/dto"
	"github.com/flintbits/drafenex-backend/internal/services"
	"github.com/gin-gonic/gin"
)

const ContextUserID = "user_id"

type OrganizerHandler struct {
	organizerService *services.OrganizerService
}

func NewOrganizerHandler(organizerService *services.OrganizerService) *OrganizerHandler {
	return &OrganizerHandler{
		organizerService: organizerService,
	}
}

func (h *OrganizerHandler) CreateOrganizer() gin.HandlerFunc {
	return func(c *gin.Context) {

		userIDValue, exists := c.Get(ContextUserID)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid user id",
			})
			return
		}

		userID, ok := userIDValue.(int64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid user id",
			})
			return
		}

		var input dto.CreateOrganizerInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		organizer, err := h.organizerService.CreateOrganizer(
			c.Request.Context(),
			userID,
			&input,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":            organizer.ID,
			"user_id":       organizer.UserID,
			"full_name":     organizer.FullName,
			"phone_number":  organizer.PhoneNumber,
			"company_name":  organizer.CompanyName,
			"avatar_url":    organizer.AvatarUrl,
			"website":       organizer.Website,
			"bio":           organizer.Bio,
			"address_line1": organizer.AddressLine1,
			"address_line2": organizer.AddressLine2,
			"city":          organizer.City,
			"state":         organizer.State,
			"country":       organizer.Country,
			"postal_code":   organizer.PostalCode,
			"status":        organizer.Status,
			"is_verified":   organizer.IsVerified,
			"created_at":    organizer.CreatedAt,
			"updated_at":    organizer.UpdatedAt,
		})
	}
}
