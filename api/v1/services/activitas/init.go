package activitas

import (
	"kaya-backend/models"
	"kaya-backend/repository"

	logger "kaya-backend/library/logger/v2"
)

// ActivitasService ..
type ActivitasService struct {
	General                *models.GeneralModel
	OttoLog                logger.KayalogInterface
	balanceTransactionRepo repository.BalanceTransactionRepository
	Database               repository.DbPostgres
}

// InitiateActivitasService ..
func InitiateActivitasService(gen *models.GeneralModel) *ActivitasService {
	return &ActivitasService{
		General:                gen,
		balanceTransactionRepo: repository.NewBalanceTransactionRepository(gen, repository.Dbcon),
	}
}
