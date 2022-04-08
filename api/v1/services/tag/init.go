package tag

import (
	"kaya-backend/models"
	"kaya-backend/repository"

	logger "kaya-backend/library/logger/v2"
)

// TagService ..
type TagService struct {
	General  *models.GeneralModel
	OttoLog  logger.KayalogInterface
	tagRepo repository.TagRepository
	Database repository.DbPostgres
}

// InitiateTagService ..
func InitiateTagService(gen *models.GeneralModel) *TagService {
	return &TagService{
		General:  gen,
		tagRepo: repository.NewTagRepository(gen, repository.Dbcon) ,
	}
}
