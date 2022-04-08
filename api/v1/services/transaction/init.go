package transaction

import (
	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"
	"kaya-backend/repository"
)

// TransactionService ..
type TransactionService struct {
	General                       *models.GeneralModel
	OttoLog                       logger.KayalogInterface
	TransactionRepo               repository.TransactionRepository
	BalanceTransactionRepo        repository.BalanceTransactionRepository
	CustomerRepo                  repository.CustomerRepository
	PropertyRepo                  repository.PropertyRepository
	CustomerPropertyLotRepo       repository.CustomerPropertyLotRepository
	CustomerPropertySecondaryRepo repository.CustomerPropertySecondaryRepository
	Database                      repository.DbPostgres
}

// InitiateTransactionInterface ..
func InitiateTransactionInterface(gen *models.GeneralModel) *TransactionService {
	return &TransactionService{
		General:                       gen,
		TransactionRepo:               repository.NewTransactionRepository(gen, repository.Dbcon),
		CustomerRepo:                  repository.NewCustomerRepository(gen, repository.Dbcon),
		PropertyRepo:                  repository.NewPropertyRepository(gen, repository.Dbcon),
		CustomerPropertyLotRepo:       repository.NewCustomerPropertyLotRepository(gen, repository.Dbcon),
		CustomerPropertySecondaryRepo: repository.NewCustomerPropertySecondaryRepository(gen, repository.Dbcon),
		BalanceTransactionRepo:        repository.NewBalanceTransactionRepository(gen, repository.Dbcon),
	}
}
