package banner

import (
	"kaya-backend/models"
	"kaya-backend/repository"

	logger "kaya-backend/library/logger/v2"
)

// BannerService ..
type BannerService struct {
	General    *models.GeneralModel
	OttoLog    logger.KayalogInterface
	bannerRepo repository.BannerRepository
	Database   repository.DbPostgres
}

// InitiatebannerService ..
func InitiatebannerService(gen *models.GeneralModel) *BannerService {
	return &BannerService{
		General:    gen,
		bannerRepo: repository.NewBannerRepository(gen, repository.Dbcon),
	}
}
