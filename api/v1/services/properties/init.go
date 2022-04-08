package properties

import (
	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"
	"kaya-backend/repository"
)

type PropertyService struct {
	General                       *models.GeneralModel
	OttoLog                       logger.KayalogInterface
	propertyRepo                  repository.PropertyRepository
	customerPropertyLotRepo       repository.CustomerPropertyLotRepository
	customerPropertySecondaryRepo repository.CustomerPropertySecondaryRepository
	customerRepo                  repository.CustomerRepository
	Database                      repository.DbPostgres
}

func InitiatePropretyInterface(gen *models.GeneralModel) *PropertyService {
	return &PropertyService{
		General:                       gen,
		propertyRepo:                  repository.NewPropertyRepository(gen, repository.Dbcon),
		customerPropertyLotRepo:       repository.NewCustomerPropertyLotRepository(gen, repository.Dbcon),
		customerRepo:                  repository.NewCustomerRepository(gen, repository.Dbcon),
		customerPropertySecondaryRepo: repository.NewCustomerPropertySecondaryRepository(gen, repository.Dbcon),
	}
}
