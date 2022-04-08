package repository

import (
	"fmt"
	"kaya-backend/models"
	"time"

	"github.com/jinzhu/gorm"
)

func NewBannerRepository(gen *models.GeneralModel, db *gorm.DB) *bannerRepository {
	return &bannerRepository{
		General: gen,
		DB:      db,
	}
}

type (
	BannerRepository interface {
		GetAll() (res models.Banner, err error)
	}
	bannerRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo bannerRepository) GetAll() (res models.Banner, err error) {
	fmt.Println(">>> bannerRepository - GetAll <<<")
	defer timeTrack(time.Now(), "GetAll")

	err = repo.DB.First(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}
