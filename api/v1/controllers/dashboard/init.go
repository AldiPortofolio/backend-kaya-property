package dashboard

import (
	"kaya-backend/models"

	"kaya-backend/api/v1/services/dashboard"

	"github.com/gin-gonic/gin"
)

type (
	dashboardController struct {
		Gen              *models.GeneralModel
		DashboardService dashboard.DashboardService
	}

	DashboardController interface {
		Dashboard(ctx *gin.Context)
	}
)

func InitiateCustomerInterface(gen *models.GeneralModel) *dashboardController {
	return &dashboardController{
		Gen:              gen,
		DashboardService: *dashboard.InitiateDashboardInterface(gen),
	}
}
