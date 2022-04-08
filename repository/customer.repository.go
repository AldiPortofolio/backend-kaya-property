package repository

import (
	"fmt"
	"kaya-backend/models"
	"time"

	"github.com/jinzhu/gorm"
)

func NewCustomerRepository(gen *models.GeneralModel, db *gorm.DB) *customerRepository {
	return &customerRepository{
		General: gen,
		DB:      db,
	}
}

type (
	CustomerRepository interface {
		Save(req models.Customers) (models.Customers, error)
		FindByID(uint64) (models.Customers, error)
		FindMe(uint64) (models.Me, error)
		WithTrx(*gorm.DB) customerRepository
		SaveGuest(req models.Guest) (models.Guest, error)
	}
	customerRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo customerRepository) WithTrx(trxHandle *gorm.DB) customerRepository {
	fmt.Println(">>> customerRepository - WithTrx <<<")
	defer timeTrack(time.Now(), "customerRepository-WithTrx")
	repo.DB = trxHandle
	return repo
}

func (repo customerRepository) UpdateBalance(req models.Customers) (models.Customers, error) {
	fmt.Println(">>> customerRepository - Save <<<")
	defer timeTrack(time.Now(), "customerRepository-Save")

	res := models.Customers{}

	err := repo.DB.Save(&req).Error
	if err != nil {
		return res, err
	}

	return req, err
}

func (repo customerRepository) Save(req models.Customers) (models.Customers, error) {
	fmt.Println(">>> customerRepository - Save <<<")
	defer timeTrack(time.Now(), "customerRepository-Save")

	res := models.Customers{}

	tx := repo.DB.Begin()

	err := tx.Save(&req).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()

	return req, err
}

func (repo customerRepository) FindByID(customerID uint64) (models.Customers, error) {
	fmt.Println(">>> customerRepository - FindByID <<<")
	defer timeTrack(time.Now(), "customerRepository-FindByID")

	res := models.Customers{}

	tx := repo.DB.Begin()

	err := Dbcon.Where("id = ?", customerID).Preload("CustomerDetails").Preload("CustomerAccounts").Preload("CustomerAccounts.Bank").Find(&res).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()

	return res, nil
}

func (repo customerRepository) FindMe(customerID uint64) (models.Me, error) {
	fmt.Println(">>> customerRepository - FindByID <<<")
	defer timeTrack(time.Now(), "customerRepository-FindByID")

	res := models.Me{}

	tx := repo.DB.Begin()

	err := Dbcon.Where("id = ?", customerID).Preload("CustomerDetails").Preload("CustomerAccounts").Preload("CustomerAccounts.Bank").Find(&res).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()

	return res, nil
}

func (repo customerRepository) SaveGuest(req models.Guest) (models.Guest, error) {
	fmt.Println(">>> customerRepository - Save <<<")
	defer timeTrack(time.Now(), "customerRepository-Save")

	res := models.Guest{}

	tx := repo.DB.Begin()

	err := tx.Save(&req).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()

	return req, err
}
