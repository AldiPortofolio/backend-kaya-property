package blog

import (
	"kaya-backend/models"
	"kaya-backend/repository"

	logger "kaya-backend/library/logger/v2"
)

// BlogService ..
type BlogService struct {
	General  *models.GeneralModel
	OttoLog  logger.KayalogInterface
	blogRepo repository.BlogRepository
	Database repository.DbPostgres
}

// InitiateBlogService ..
func InitiateBlogService(gen *models.GeneralModel) *BlogService {
	return &BlogService{
		General:  gen,
		blogRepo: repository.NewBlogRepository(gen, repository.Dbcon) ,
	}
}
