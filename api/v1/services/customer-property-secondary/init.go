package customerpropertysecondary

import (
	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"
	"kaya-backend/repository"
)

type CustomerPropertySecondaryService struct {
	General                       *models.GeneralModel
	OttoLog                       logger.KayalogInterface
	CustomerPropertySecondaryRepo repository.CustomerPropertySecondaryRepository
	PropertyRepo                  repository.PropertyRepository
	Database                      repository.DbPostgres
}

func InitiateCustomerPropertySecondaryInterface(gen *models.GeneralModel) *CustomerPropertySecondaryService {
	return &CustomerPropertySecondaryService{
		General:                       gen,
		CustomerPropertySecondaryRepo: repository.NewCustomerPropertySecondaryRepository(gen, repository.Dbcon),
		PropertyRepo:                  repository.NewPropertyRepository(gen, repository.Dbcon),
	}
}
