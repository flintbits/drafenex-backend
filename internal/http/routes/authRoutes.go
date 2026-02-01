package routes

import (
	"github.com/flintbits/drafenex-backend/internal/config"
	"github.com/flintbits/drafenex-backend/internal/http/handlers"
	"github.com/flintbits/drafenex-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(cfg *config.Config, router *gin.Engine, UserService *services.UserService, AuthService *services.AuthService) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", handlers.CreateUserHandler(UserService))
		auth.POST("/login", handlers.LoginHandler(AuthService))
		auth.POST("/logout", handlers.LogoutHandler())
		auth.GET("/me", handlers.Me(cfg, UserService))
	}
}
