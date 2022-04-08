package repository

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"time"

	"github.com/jinzhu/gorm"
)

func NewBalanceTransactionRepository(gen *models.GeneralModel, db *gorm.DB) *balanceTransactionRepository {
	return &balanceTransactionRepository{
		General: gen,
		DB:      db,
	}
}

type (
	BalanceTransactionRepository interface {
		HistoryTopup(req request.ActivitasFilter) ([]models.BalanceTransaction, int, error)
		HistoryTransaction(req request.ActivitasFilter) ([]models.Transactions, int, error)
		Save(req models.BalanceTransaction) (models.BalanceTransaction, error)
		WithTrx(*gorm.DB) balanceTransactionRepository
	}
	balanceTransactionRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo balanceTransactionRepository) WithTrx(trxHandle *gorm.DB) balanceTransactionRepository {
	fmt.Println(">>> balanceTransactionRepository - WithTrx <<<")
	defer timeTrack(time.Now(), "balanceTransactionRepository-WithTrx")
	repo.DB = trxHandle
	return repo
}

func (repo balanceTransactionRepository) Save(req models.BalanceTransaction) (models.BalanceTransaction, error) {
	fmt.Println(">>> transactionRepository - Save <<<")
	defer timeTrack(time.Now(), "Save")

	err := repo.DB.Save(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

func (repo balanceTransactionRepository) HistoryTopup(req request.ActivitasFilter) (res []models.BalanceTransaction, total int, err error) {
	fmt.Println(">>> propertyRepository - GetAll <<<")
	defer timeTrack(time.Now(), "propertyRepository-GetAll")

	db := repo.DB

	transactionType := []string{"TOPUP", "WITHDRAWAL", "PENJUALAN_PROPERTY", "PEMBELIAN_LOT", "PENJUALAN_LOT"}
	db = db.Where(`transaction_type in (?)`, transactionType)

	now := time.Now()
	y, m, d := now.Date()
	date := fmt.Sprintf("%d-%d-%d", y, int(m), d)

	if req.Filter != "" {

		if req.Filter == "today" {
			db = db.Where(`date(created_at) = date('` + date + `')`)
		}

		if req.Filter == "month" {
			db = db.Where(`TO_CHAR(DATE(created_at), 'Month') = TO_CHAR(DATE('` + date + `'), 'Month')`)
		}

		if req.Filter == "year" {
			db = db.Where(`date_part('year', created_at) = date_part('year', timestamp '` + date + `')`)
		}
	}

	if req.Filter == "" {
		db = db.Where(`date(created_at) = date('` + date + `')`)
	}

	db = db.Where(`customer_id = ?`, req.CustomerID)

	if err := db.Order("id DESC").Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		return res, total, err
	}
	return res, total, nil
}

func (repo balanceTransactionRepository) HistoryTransaction(req request.ActivitasFilter) (res []models.Transactions, total int, err error) {
	fmt.Println(">>> propertyRepository - GetAll <<<")
	defer timeTrack(time.Now(), "propertyRepository-GetAll")

	db := repo.DB

	now := time.Now()
	y, m, d := now.Date()
	date := fmt.Sprintf("%d-%d-%d", y, int(m), d)

	if req.Filter != "" {
		if req.Filter == "today" {
			db = db.Where(`date(created_at) = date('` + date + `')`)
		}

		if req.Filter == "month" {
			db = db.Where(`TO_CHAR(DATE(created_at), 'Month') = TO_CHAR(DATE('` + date + `'), 'Month')`)
		}

		if req.Filter == "year" {
			db = db.Where(`date_part('year', created_at) = date_part('year', timestamp '` + date + `')`)
		}
	}

	if req.Filter == "" {
		db = db.Where(`date(created_at) = date('` + date + `')`)
	}

	db = db.Where(`customer_id = ?`, req.CustomerID)

	if err := db.Order("id DESC").Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Preload("PaymentMethod").Preload("TransactionDetail.Property").Preload("Status").Find(&res).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		return res, total, err
	}
	return res, total, nil
}
