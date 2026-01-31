package routes

import (
	"github.com/daixiang0/gci/pkg/config"
	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterAuthRoutes(router *gin.Engine, pool *pgxpool.Pool, cfg *config.Config) {
	auth := router.Group("/auth")
	{
		auth.POST("/register")
		auth.POST("/login")
	}
}
