package repository

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// NewCustomerPropertySecondaryRepository ..
func NewCustomerPropertySecondaryRepository(gen *models.GeneralModel, db *gorm.DB) *customerPropertySecondaryRepository {
	return &customerPropertySecondaryRepository{
		General: gen,
		DB:      db,
	}
}

// CustomerPropertySecondaryRepository ..
type (
	CustomerPropertySecondaryRepository interface {
		Save(models.CustomerPropertySecondaries) (models.CustomerPropertySecondaries, error)
		GetAll(req request.FilterProperty) ([]models.CustomerPropertySecondaries, int, error)
		FindByID(int) (res models.CustomerPropertySecondaries, err error)
		FindBySlug(string) (res []models.CustomerPropertySecondaries, err error)
		FindBySlugClosed(string) (res []models.CustomerPropertySecondaries, err error)
		WithTrx(*gorm.DB) customerPropertySecondaryRepository
		Filter(req request.FilterPropertySecondary) (res []models.CustomerPropertySecondaries, total int, err error)
		SumLotPrice(req request.FilterPropertySecondary) (res models.ResSumCustomerPropertySecondaries, err error)
		ListByCustomer(req request.FilterPropertySecondary) (res []models.CustomerPropertySecondaries, err error)
		Delete(req models.CustomerPropertySecondaries) error
	}
	customerPropertySecondaryRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo customerPropertySecondaryRepository) SumLotPrice(req request.FilterPropertySecondary) (res models.ResSumCustomerPropertySecondaries, err error) {
	fmt.Println(">>> Database - GetPropertySecondaries <<<")
	defer timeTrack(time.Now(), "SumLotPricePropertySecondaries")
	db := repo.DB
	id := strconv.Itoa(req.CustomerID)
	var query string

	if req.Status != "" && req.CustomerID != 0 {
		query = "where status = '" + req.Status + "' and customer_id = " + id
	} else if req.Status == "" && req.CustomerID != 0 {
		query = "where customer_id = " + id
	} else if req.Status != "" && req.CustomerID == 0 {
		query = "where status = '" + req.Status + "'"
	}

	if err := db.Raw("select sum(lot) as total_lot, sum(price_per_lot) as total_price from customer_property_secondaries " + query).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (repo customerPropertySecondaryRepository) Filter(req request.FilterPropertySecondary) (res []models.CustomerPropertySecondaries, total int, err error) {
	fmt.Println(">>> Database - GetPropertySecondaries <<<")
	defer timeTrack(time.Now(), "GetFilterPropertySecondaries")

	db := repo.DB
	db = db.Preload("Property.PropertyFee").Preload("Property.PropertyPhotos").Preload("Property.City").Preload("Property.City.Province")

	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	if req.CustomerID != 0 {
		db = db.Where("customer_id = ?", req.CustomerID)
	}

	if err := db.Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		return res, total, err
	}

	return res, total, nil
}

func (repo customerPropertySecondaryRepository) ListByCustomer(req request.FilterPropertySecondary) (res []models.CustomerPropertySecondaries, err error) {
	fmt.Println(">>> Database - GetPropertySecondaries <<<")
	defer timeTrack(time.Now(), "GetFilterPropertySecondaries")

	db := repo.DB
	db = db.Preload("Property.PropertyFee").Preload("Property.PropertyPhotos").Preload("Property.City").Preload("Property.City.Province")

	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	if req.CustomerID != 0 {
		db = db.Where("customer_id = ?", req.CustomerID)
	}

	if err := db.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (repo customerPropertySecondaryRepository) WithTrx(trxHandle *gorm.DB) customerPropertySecondaryRepository {
	fmt.Println(">>> CustomerPropertySecondaryRepository - WithTrx <<<")
	defer timeTrack(time.Now(), "CustomerPropertySecondaryRepository-WithTrx")
	repo.DB = trxHandle
	return repo
}

func (repo customerPropertySecondaryRepository) Save(req models.CustomerPropertySecondaries) (models.CustomerPropertySecondaries, error) {
	fmt.Println(">>> CustomerPropertySecondaryRepository - Save <<<")
	defer timeTrack(time.Now(), "Save")

	err := repo.DB.Omit("CustomerPropertyLot").Save(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

func (repo customerPropertySecondaryRepository) Delete(req models.CustomerPropertySecondaries) error {
	fmt.Println(">>> CustomerPropertySecondaryRepository - Delete <<<")
	defer timeTrack(time.Now(), "Save")

	err := repo.DB.Save(&req).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo customerPropertySecondaryRepository) GetAll(req request.FilterProperty) (res []models.CustomerPropertySecondaries, total int, err error) {
	fmt.Println(">>> Database - GetPropertySecondaries <<<")
	defer timeTrack(time.Now(), "GetPropertySecondaries")

	db := repo.DB
	db = db.Preload("Property.PropertyFee").Preload("Property.PropertyPhotos").Preload("Property.City").Preload("Property.City.Province")
	db = db.Joins("INNER JOIN properties ON properties.id = customer_property_secondaries.property_id")
	db = db.Select("customer_property_secondaries.property_id")

	if req.Name != "" {
		db = db.Where(`properties.name ilike '%` + req.Name + `%'`)
	}

	if req.CityID != 0 {
		db = db.Where("properties.city_id = ?", req.CityID)
	}

	if req.CityID != 0 {
		db = db.Where("properties.city_id = ?", req.CityID)
	}

	if req.CustomerID != 0 {
		db = db.Where("customer_property_secondaries.customer_id = ?", req.CustomerID)
	}

	db = db.Where("properties.is_sold = ?", false)
	db = db.Where("customer_property_secondaries.status = ?", "OPEN")

	if err := db.Group("property_id").Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		return res, total, err
	}

	return res, total, nil
}

func (repo customerPropertySecondaryRepository) FindByID(ID int) (res models.CustomerPropertySecondaries, err error) {
	fmt.Println(">>> customerPropertySecondaryRepository - FindByID <<<")
	defer timeTrack(time.Now(), "customerPropertySecondaryRepository-FindByID")

	db := repo.DB
	db = db.Preload("CustomerPropertyLot").Preload("Property").Preload("Property.PropertyPhotos").Preload("Property.PropertyFee")

	if err := db.Where("id = ?", ID).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (repo customerPropertySecondaryRepository) FindBySlug(slug string) (res []models.CustomerPropertySecondaries, err error) {
	fmt.Println(">>> Database - PropertiesSecondaryBySlug <<<")
	defer timeTrack(time.Now(), "PropertiesSecondaryBySlug")

	db := repo.DB
	db = db.Preload("Property.PropertyFee").Preload("Property.PropertyPhotos").Preload("Property.City").Preload("Property.City.Province").Preload("Customer")
	db = db.Joins("INNER JOIN properties ON properties.id = customer_property_secondaries.property_id")

	if err := db.Where("properties.slug = ?", slug).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (repo customerPropertySecondaryRepository) FindBySlugClosed(slug string) (res []models.CustomerPropertySecondaries, err error) {
	fmt.Println(">>> Database - PropertiesSecondaryBySlug <<<")
	defer timeTrack(time.Now(), "PropertiesSecondaryBySlug")

	db := repo.DB
	db = db.Preload("Property.PropertyFee").Preload("Property.PropertyPhotos").Preload("Property.City").Preload("Property.City.Province").Preload("Customer")
	db = db.Joins("INNER JOIN properties ON properties.id = customer_property_secondaries.property_id")
	db = db.Where("customer_property_secondaries.status = ?", "OPEN")
	if err := db.Where("properties.slug = ?", slug).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
