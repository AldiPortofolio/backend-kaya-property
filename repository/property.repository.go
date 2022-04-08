package repository

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"time"

	"github.com/jinzhu/gorm"
)

// NewPropertyRepository ..
func NewPropertyRepository(gen *models.GeneralModel, db *gorm.DB) *propertyRepository {
	return &propertyRepository{
		General: gen,
		DB:      db,
	}
}

// PropertyRepository ..
type (
	PropertyRepository interface {
		Save(models.Properties) (models.Properties, error)
		GetAll(req request.FilterProperty) ([]models.Properties, int, error)
		GetAllPortfolio(req request.FilterProperty) ([]models.Properties, int, error)
		FindBySlug(string) (res models.DetailProperties, err error)
		FindByID(int) (res models.DetailProperties, err error)
		Find(int) (res models.Properties, err error)
		WithTrx(*gorm.DB) propertyRepository
	}
	propertyRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo propertyRepository) WithTrx(trxHandle *gorm.DB) propertyRepository {
	fmt.Println(">>> propertyRepository - WithTrx <<<")
	defer timeTrack(time.Now(), "propertyRepository-WithTrx")
	repo.DB = trxHandle
	return repo
}

func (repo propertyRepository) Save(req models.Properties) (models.Properties, error) {
	fmt.Println(">>> propertyRepository - Save <<<")
	defer timeTrack(time.Now(), "Save")

	err := repo.DB.Save(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

func (repo propertyRepository) GetAll(req request.FilterProperty) (res []models.Properties, total int, err error) {
	fmt.Println(">>> propertyRepository - GetAll <<<")
	defer timeTrack(time.Now(), "propertyRepository-GetAll")

	db := repo.DB
	db = db.Preload("PropertyFee").Preload("PropertyPhotos").Preload("City").Preload("City.Province")

	if req.Name != "" {
		db = db.Where(`name ilike '%` + req.Name + `%'`)
	}

	if req.CityID != 0 {
		db = db.Where("city_id = ?", req.CityID)
	}
	db = db.Where("is_sold = ?", false)

	if err := db.Order("id DESC").Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		return res, total, err
	}
	return res, total, nil
}

func (repo propertyRepository) GetAllPortfolio(req request.FilterProperty) (res []models.Properties, total int, err error) {
	fmt.Println(">>> propertyRepository - GetAll <<<")
	defer timeTrack(time.Now(), "propertyRepository-GetAll")

	db := repo.DB
	db = db.Preload("PropertyFee").Preload("PropertyPhotos").Preload("City").Preload("City.Province")

	if req.Name != "" {
		db = db.Where(`name ilike '%` + req.Name + `%'`)
	}

	if req.CityID != 0 {
		db = db.Where("city_id = ?", req.CityID)
	}

	db = db.Where("sold_price != 0 and sold_date != '0001-01-01'")

	if err := db.Order("id DESC").Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		return res, total, err
	}
	return res, total, nil
}

func (repo propertyRepository) FindBySlug(slug string) (res models.DetailProperties, err error) {
	fmt.Println(">>> Database - GetDetailProperty <<<")
	defer timeTrack(time.Now(), "GetDetailProperty")

	db := repo.DB
	db = db.Preload("PropertyPhotos").Preload("PropertyFee").Preload("City").Preload("City.Province")

	if err := db.Where("slug = ?", slug).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (repo propertyRepository) FindByID(ID int) (res models.DetailProperties, err error) {
	fmt.Println(">>> Database - GetPropertyByID <<<")
	defer timeTrack(time.Now(), "GetPropertyByID")

	db := repo.DB
	db = db.Preload("PropertyFee").Preload("PropertyPhotos")

	if err := db.Where("id = ?", ID).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (repo propertyRepository) Find(ID int) (res models.Properties, err error) {
	fmt.Println(">>> Database - GetPropertyByID <<<")
	defer timeTrack(time.Now(), "GetPropertyByID")

	db := repo.DB
	db = db.Preload("PropertyFee").Preload("PropertyPhotos")

	if err := db.Where("id = ?", ID).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
