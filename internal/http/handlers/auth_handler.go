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

		http.SetCookie(
			c.Writer, &http.Cookie{
				Name:     "access_token",
				Value:    token,
				HttpOnly: true,
				Secure:   false, //true in production
				SameSite: http.SameSiteLaxMode,
				Path:     "/",
				MaxAge:   60 * 15,
			})

		c.JSON(http.StatusOK, gin.H{
			// "token": token,
			"user": gin.H{
				"id":           user.ID,
				"email":        user.Email,
				"role":         user.Role,
				"is_onboarded": user.IsOnboarded,
			},
		})
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "access_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   -1,
		})

		c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
	}
}

//ToDo: Secure:   false, //true in production
