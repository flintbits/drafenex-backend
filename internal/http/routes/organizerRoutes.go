package routes

import (
	"github.com/flintbits/drafenex-backend/internal/config"
	"github.com/flintbits/drafenex-backend/internal/http/handlers"
	"github.com/flintbits/drafenex-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterOrganizerRoutes(cfg *config.Config, router *gin.RouterGroup, organizerService *services.OrganizerService) {
	organizerHandler := handlers.NewOrganizerHandler(organizerService)

	organizer := router.Group("/organizers")
	{
		organizer.POST("/onboarding", organizerHandler.CreateOrganizer())
	}
}
