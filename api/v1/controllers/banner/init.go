package banner

import (
	"kaya-backend/models"

	"kaya-backend/api/v1/services/banner"

	"github.com/gin-gonic/gin"
)

type (
	bannerontroller struct {
		Gen           *models.GeneralModel
		BannerService banner.BannerService
	}

	BannerController interface {
		Banner(ctx *gin.Context)
	}
)

func InitiateBannerInterface(gen *models.GeneralModel) *bannerontroller {
	return &bannerontroller{
		Gen:           gen,
		BannerService: *banner.InitiatebannerService(gen),
	}
}
