package repository

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"time"

	"github.com/jinzhu/gorm"
)

// NewCustomerPropertyLotRepository ..
func NewCustomerPropertyLotRepository(gen *models.GeneralModel, db *gorm.DB) *customerPropertyLotRepository {
	return &customerPropertyLotRepository{
		General: gen,
		DB:      db,
	}
}

// CustomerPropertyLotRepository ..
type (
	CustomerPropertyLotRepository interface {
		Save(models.CustomerPropertyLots) (models.CustomerPropertyLots, error)
		GetAll(req request.FilterProperty) ([]models.CustomerPropertyLots, int, error)
		GetAllMe(req request.FilterProperty) (res []models.CustomerPropertyLots, total int, err error)
		FindByPropertyID(int) (res []models.CustomerPropertyLots, err error)
		FindByCustomerID(int) (res []models.CustomerPropertyLots, err error)
		FindByCustomerIDAndPropertyId(int, int) (res models.CustomerPropertyLots, err error)
		FindByFilterProperty(request.FilterPortfolio) (res []models.CustomerPropertyLots, err error)
		FindByID(int) (res models.CustomerPropertyLots, err error)
		WithTrx(*gorm.DB) customerPropertyLotRepository
		UpdateCustomerPropertyLot(req models.CustomerPropertyLots) error
	}
	customerPropertyLotRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

// WithTrx
func (repo customerPropertyLotRepository) WithTrx(trxHandle *gorm.DB) customerPropertyLotRepository {
	fmt.Println(">>> customerPropertyLotRepository - WithTrx <<<")
	defer timeTrack(time.Now(), "customerPropertyLotRepository-WithTrx")
	repo.DB = trxHandle
	return repo
}

// Save
func (repo customerPropertyLotRepository) Save(req models.CustomerPropertyLots) (models.CustomerPropertyLots, error) {
	fmt.Println(">>> customerPropertyLotRepository - Save <<<")
	defer timeTrack(time.Now(), "customerPropertyLotRepository-Save")

	err := repo.DB.Save(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

//GetAll
func (repo customerPropertyLotRepository) GetAll(req request.FilterProperty) (res []models.CustomerPropertyLots, total int, err error) {
	fmt.Println(">>> Database - GetPropertySecondaries <<<")
	defer timeTrack(time.Now(), "GetPropertySecondaries")

	db := repo.DB
	db = db.Preload("Property.PropertyFee").Preload("Property.PropertyPhotos").Preload("Property.City").Preload("Property.City.Province")

	if req.CustomerID != 0 {
		db = db.Where("customer_id = ?", req.CustomerID)
	}

	if err := db.Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		return res, total, err
	}

	return res, total, nil
}

//GetAllMe
func (repo customerPropertyLotRepository) GetAllMe(req request.FilterProperty) (res []models.CustomerPropertyLots, total int, err error) {
	fmt.Println(">>> Database - GetPropertySecondaries <<<")
	defer timeTrack(time.Now(), "GetPropertySecondaries")

	db := repo.DB
	db = db.Preload("Property.PropertyFee").Preload("Property.PropertyPhotos").Preload("Property.City").Preload("Property.City.Province")
	db = db.Joins("INNER JOIN properties ON properties.id = customer_property_lots.property_id")
	db = db.Where("customer_id = ?", req.CustomerID)

	if req.Status == "OPEN" {
		db = db.Where("properties.is_sold = ?", false)
		db = db.Where("properties.presentase != 100")
	}

	if req.Status == "SOLD" {
		db = db.Where("properties.is_sold = ?", true)
	}

	if req.Status == "DONE" {
		db = db.Where("properties.is_sold = ?", false)
		db = db.Where("properties.presentase = 100")
	}

	if err := db.Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		return res, total, err
	}

	return res, total, nil
}

// FindByPropertyID ..
func (repo customerPropertyLotRepository) FindByPropertyID(ID int) (res []models.CustomerPropertyLots, err error) {
	fmt.Println(">>> Database - FindByPropertyID <<<")
	defer timeTrack(time.Now(), "FindByPropertyID")

	db := repo.DB
	if err := db.Where("property_id = ?", ID).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

// FindByCustomerID ..
func (repo customerPropertyLotRepository) FindByCustomerID(ID int) (res []models.CustomerPropertyLots, err error) {
	fmt.Println(">>> Database - FindByPropertyID <<<")
	defer timeTrack(time.Now(), "FindByPropertyID")

	db := repo.DB
	if err := db.Where("customer_id = ?", ID).Preload("Property").Preload("Property.PropertyFee").Preload("Property.PropertyPhotos").Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

// FindByCustomerID ..
func (repo customerPropertyLotRepository) FindByCustomerIDAndPropertyId(CustomerID int, propertyId int) (res models.CustomerPropertyLots, err error) {
	fmt.Println(">>> Database - FindByPropertyVtCustomerID <<<")
	defer timeTrack(time.Now(), "FindByPropertyID")

	db := repo.DB
	if err := db.Where("customer_id = ?", CustomerID).Where("property_id = ?", propertyId).Preload("Property").Preload("Property.PropertyFee").Preload("Property.PropertyPhotos").First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

// FindByFilter
func (repo customerPropertyLotRepository) FindByFilterProperty(req request.FilterPortfolio) (res []models.CustomerPropertyLots, err error) {
	fmt.Println(">>> Database - FindByFilter <<<")
	defer timeTrack(time.Now(), "customerPropertyLotRepository - FindByFilter")

	db := repo.DB
	db = db.Preload("Property").Preload("Property.PropertyFee")
	db = db.Joins("INNER JOIN properties ON properties.id = customer_property_lots.property_id")

	if req.Status == "OPEN" {
		db = db.Where("properties.is_sold = ?", false)
		db = db.Where("properties.presentase != 100")
	}

	if req.Status == "SOLD" {
		db = db.Where("properties.is_sold = ?", true)
	}

	if req.Status == "DONE" {
		db = db.Where("properties.is_sold = ?", false)
		db = db.Where("properties.presentase = 100")
	}

	if req.CustomerID != 0 {
		db = db.Where("customer_property_lots.customer_id = ?", req.CustomerID)
	}

	if err := db.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

// FindByPropertyID ..
func (repo customerPropertyLotRepository) FindByID(ID int) (res models.CustomerPropertyLots, err error) {
	fmt.Println(">>> Database - FindByPropertyID <<<")
	defer timeTrack(time.Now(), "FindByPropertyID")

	db := repo.DB
	db = db.Preload("Property.PropertyFee").Preload("Property.PropertyPhotos").Preload("Property.City").Preload("Property.City.Province")
	if err := db.Where("id = ?", ID).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil

}

func (repo customerPropertyLotRepository) UpdateCustomerPropertyLot(req models.CustomerPropertyLots) (err error) {
	fmt.Println(">>> Database - findByCustomerAndProperty <<<")
	defer timeTrack(time.Now(), "FindByPropertyID")

	db := repo.DB

	err = db.Save(&req).Error
	if err != nil {
		return err
	}
	return nil
}
