package models

import "time"

type Ticket struct {
	ID             int             `gorm:"column:id" json:"id"`
	CustomerID     int             `gorm:"column:customer_id" json:"customer_id"`
	UserID         int             `gorm:"column:user_id" json:"user_id"`
	Title          string          `gorm:"column:title" json:"title"`
	Content        string          `gorm:"column:content" json:"content"`
	CreatedAt      time.Time       `gorm:"column:created_at" json:"created_at"`
	TicketComments []TicketComment `gorm:"references:ticket_id" json:"ticket_comments"`
	Customer       Customers       `gorm:"references:CustomerID" json:"customer"`
	User           User            `gorm:"references:UserID" json:"user"`
	Status         string          `gorm:"column:status" json:"status"`
}

func (p *Ticket) TableName() string {
	return "tickets"
}
