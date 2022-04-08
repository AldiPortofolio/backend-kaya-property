package activitas

import (
	"kaya-backend/models"

	"kaya-backend/api/v1/services/activitas"

	"github.com/gin-gonic/gin"
)

type (
	activitasController struct {
		Gen              *models.GeneralModel
		ActivitasService activitas.ActivitasService
	}

	ActivitasController interface {
		Activitas(ctx *gin.Context)
	}
)

func InitiateActivitasInterface(gen *models.GeneralModel) *activitasController {
	return &activitasController{
		Gen:              gen,
		ActivitasService: *activitas.InitiateActivitasService(gen),
	}
}
