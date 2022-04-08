package repository

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"time"

	"github.com/jinzhu/gorm"
)

// NewTicketRepository ..
func NewTicketRepository(gen *models.GeneralModel, db *gorm.DB) *ticketRepository {
	return &ticketRepository{
		General: gen,
		DB:      db,
	}
}

// TicketRepository ..
type (
	TicketRepository interface {
		Save(models.Ticket) (models.Ticket, error)
		SaveTicketComment(models.TicketComment) (models.TicketComment, error)
		GetAll(req request.FilterTicket) ([]models.Ticket, int, error)
		Detail(int) (models.Ticket, error)
		WithTrx(*gorm.DB) ticketRepository
	}
	ticketRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

// WithTrx
func (repo ticketRepository) WithTrx(trxHandle *gorm.DB) ticketRepository {
	fmt.Println(">>> ticketRepository - WithTrx <<<")
	defer timeTrack(time.Now(), "ticketRepository-WithTrx")
	repo.DB = trxHandle
	return repo
}

// Save
func (repo ticketRepository) Save(req models.Ticket) (models.Ticket, error) {
	fmt.Println(">>> ticketRepository - Save <<<")
	defer timeTrack(time.Now(), "ticketRepository-Save")

	err := repo.DB.Save(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

// SaveTicketComment
func (repo ticketRepository) SaveTicketComment(req models.TicketComment) (models.TicketComment, error) {
	fmt.Println(">>> ticketRepository - SaveTicketComment <<<")
	defer timeTrack(time.Now(), "ticketRepository-SaveTicketComment")

	req.IsAdmin = false;
	err := repo.DB.Save(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

// Get all
func (repo ticketRepository) GetAll(req request.FilterTicket) (res []models.Ticket, total int, err error) {
	fmt.Println(">>> ticketRepository - GetAll <<<")
	defer timeTrack(time.Now(), "ticketRepository-GetAll")

	db := repo.DB
	db = db.Preload("TicketComments").Preload("Customer").Preload("Customer.CustomerDetails").Preload("User")

	db = db.Where(`customer_id = ?`, req.CustomerID)

	if err := db.Order("id DESC").Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error; err != nil {
		return res, total, err
	}
	return res, total, nil
}

// Detail
func (repo ticketRepository) Detail(ID int) (res models.Ticket, err error) {
	fmt.Println(">>> ticketRepository - GetAll <<<")
	defer timeTrack(time.Now(), "ticketRepository-GetAll")

	db := repo.DB
	db = db.Preload("TicketComments").Preload("Customer").Preload("User")

	db = db.Where(`id = ?`, ID)

	if err := db.Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

