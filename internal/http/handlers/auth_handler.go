package handlers

import (
	"net/http"

	"github.com/flintbits/drafenex-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, user, err := authService.Login(
			c.Request.Context(),
			req.Email,
			req.Password,
		)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid credentials 3",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user": gin.H{
				"id":           user.ID,
				"email":        user.Email,
				"role":         user.Role,
				"is_onboarded": user.IsOnboarded,
			},
		})
	}
}
