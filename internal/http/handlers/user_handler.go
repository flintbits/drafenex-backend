package handlers

import (
	"net/http"

	"github.com/flintbits/drafenex-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func CreateUserHandler(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "json " + err.Error()})
			return
		}

		user, err := userService.CreateUser(
			c.Request.Context(), req.Email, req.Password, req.Role,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":           user.ID,
			"email":        user.Email,
			"role":         user.Role,
			"is_onboarded": user.IsOnboarded,
			"created_at":   user.CreatedAt,
			"updated_at":   user.UpdatedAt,
		})
	}
}
