package http

import (
	"log"

	"github.com/flintbits/drafenex-backend/internal/config"
	"github.com/flintbits/drafenex-backend/internal/container"
	"github.com/flintbits/drafenex-backend/internal/http/routes"
	"github.com/flintbits/drafenex-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config, serviceContainer *container.Container) *gin.Engine {
	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal("failed to set trusted proxies:", err)
	}

	router.Use(middleware.CORSMiddleware())

	api := router.Group("/api")

	routes.RegisterHealthRoutes(router)

	routes.RegisterAuthRoutes(cfg, api, serviceContainer.UserService, serviceContainer.AuthService)

	routes.RegisterOrganizerRoutes(cfg, api, serviceContainer.OrganizerService)

	return router
}
