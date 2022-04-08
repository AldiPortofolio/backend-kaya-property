package repository

import (
	"fmt"
	"kaya-backend/models"
	"time"

	"github.com/jinzhu/gorm"
)

// NewTitipJualRepository ..
func NewTitipJualRepository(gen *models.GeneralModel, db *gorm.DB) *titipJualRepository {
	return &titipJualRepository{
		General: gen,
		DB:      db,
	}
}

// TitipJualRepository ..
type (
	TitipJualRepository interface {
		Save(models.TitipJual) (models.TitipJual, error)
	}

	titipJualRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo titipJualRepository) Save(req models.TitipJual) (models.TitipJual, error) {
	fmt.Println(">>> titipJualRepository - Save <<<")
	defer timeTrack(time.Now(), "Save")

	db := repo.DB.Begin()
	err := db.Save(&req).Error
	if err != nil {
		db.Rollback()
		return req, err
	}
	db.Commit()

	return req, nil
}
