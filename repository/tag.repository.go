package repository

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"time"

	"github.com/jinzhu/gorm"
)

// NewTagRepository ..
func NewTagRepository(gen *models.GeneralModel, db *gorm.DB) *tagRepository {
	return &tagRepository{
		General: gen,
		DB:      db,
	}
}

// TagRepository ..
type (
	TagRepository interface {
		Filter(request.TagBlog) ([]models.Tag, error)
		WithTrx(*gorm.DB) tagRepository
	}
	tagRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo tagRepository) WithTrx(trxHandle *gorm.DB) tagRepository {
	fmt.Println(">>> tagRepository - WithTrx <<<")
	defer timeTrack(time.Now(), "tagRepository-WithTrx")
	repo.DB = trxHandle
	return repo
}

func (repo tagRepository) Filter(req request.TagBlog) ([]models.Tag, error) {
	fmt.Println(">>> tagRepository - Filter <<<")
	defer timeTrack(time.Now(), "Filter")

	var res []models.Tag
	db := repo.DB

	if len(req.Id) > 0 {
		db.Where("id in ?", req.Id)
	}

	err := db.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
