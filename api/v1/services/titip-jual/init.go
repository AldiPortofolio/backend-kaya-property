package titipjual

import (
	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"
	"kaya-backend/repository"
)

// TitipJualService ..
type TitipJualService struct {
	General       *models.GeneralModel
	OttoLog       logger.KayalogInterface
	titipJualRepo repository.TitipJualRepository
	Database      repository.DbPostgres
}

// InitiateTitipJualInterface ..
func InitiateTitipJualInterface(gen *models.GeneralModel) *TitipJualService {
	return &TitipJualService{
		General:       gen,
		titipJualRepo: repository.NewTitipJualRepository(gen, repository.Dbcon),
	}
}
