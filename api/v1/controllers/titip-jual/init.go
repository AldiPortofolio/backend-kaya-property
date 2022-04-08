package titipjual

import (
	"kaya-backend/models"

	titipjual "kaya-backend/api/v1/services/titip-jual"

	"github.com/gin-gonic/gin"
)

type (
	titipJualController struct {
		Gen              *models.GeneralModel
		TitipJualService titipjual.TitipJualService
	}

	TitipJualController interface {
		TitipJual(ctx *gin.Context)
	}
)

func InitiateTitipJualInterface(gen *models.GeneralModel) *titipJualController {
	return &titipJualController{
		Gen:              gen,
		TitipJualService: *titipjual.InitiateTitipJualInterface(gen),
	}
}
