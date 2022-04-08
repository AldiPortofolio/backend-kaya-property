package dashboard

import (
	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"
	"kaya-backend/repository"
)

// DashboardService ..
type DashboardService struct {
	General                       *models.GeneralModel
	OttoLog                       logger.KayalogInterface
	TransactionRepo               repository.TransactionRepository
	BalanceTransactionRepo        repository.BalanceTransactionRepository
	CustomerRepo                  repository.CustomerRepository
	PropertyRepo                  repository.PropertyRepository
	CustomerPropertyLotRepo       repository.CustomerPropertyLotRepository
	CustomerPropertySecondaryRepo repository.CustomerPropertySecondaryRepository
	ticketRepo                    repository.TicketRepository
	Database                      repository.DbPostgres
}

// InitiateTransactionInterface ..
func InitiateDashboardInterface(gen *models.GeneralModel) *DashboardService {
	return &DashboardService{
		General:                       gen,
		TransactionRepo:               repository.NewTransactionRepository(gen, repository.Dbcon),
		CustomerRepo:                  repository.NewCustomerRepository(gen, repository.Dbcon),
		PropertyRepo:                  repository.NewPropertyRepository(gen, repository.Dbcon),
		CustomerPropertyLotRepo:       repository.NewCustomerPropertyLotRepository(gen, repository.Dbcon),
		CustomerPropertySecondaryRepo: repository.NewCustomerPropertySecondaryRepository(gen, repository.Dbcon),
		BalanceTransactionRepo:        repository.NewBalanceTransactionRepository(gen, repository.Dbcon),
		ticketRepo:                    repository.NewTicketRepository(gen, repository.Dbcon),
	}
}
