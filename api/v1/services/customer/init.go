package customer

import (
	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"
	"kaya-backend/repository"
)

type CustomerService struct {
	General             *models.GeneralModel
	OttoLog             logger.KayalogInterface
	CustomerRepo        repository.CustomerRepository
	MembershipLevelRepo repository.MembershipLevelRepository
	Database            repository.DbPostgres
}

func InitiateCustomerInterface(gen *models.GeneralModel) *CustomerService {
	return &CustomerService{
		General:             gen,
		CustomerRepo:        repository.NewCustomerRepository(gen, repository.Dbcon),
		MembershipLevelRepo: repository.NewMembershipLevelRepository(gen, repository.Dbcon),
	}
}
