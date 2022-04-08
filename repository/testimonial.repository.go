package repository

import (
	"fmt"
	"kaya-backend/models"
	"time"

	"github.com/jinzhu/gorm"
)

// NewTestimonialsRepository ..
func NewTestimonialsRepository(gen *models.GeneralModel, db *gorm.DB) *testimonialsRepository {
	return &testimonialsRepository{
		General: gen,
		DB:      db,
	}
}

// TransactionRepository ..
type (
	TestimonialsRepository interface {
		GetAll() (res []models.Testimonials, err error)
	}
	testimonialsRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)


func (repo testimonialsRepository) GetAll() (res []models.Testimonials, err error) {
	fmt.Println(">>> testimonialsRepository - GetAll <<<")
	defer timeTrack(time.Now(), "GetAll")

	db := repo.DB
	db = db.Where("publish = true").Order("created_at DESC")

	err = db.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
