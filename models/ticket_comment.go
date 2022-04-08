package models

import "time"

type TicketComment struct {
	ID       int    `gorm:"column:id" json:"id"`
	TicketID int    `gorm:"column:ticket_id" json:"ticket_id"`
	Content  string `gorm:"column:content" json:"content"`
	IsAdmin  bool `gorm:"column:is_admin" json:"is_admin"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (p *TicketComment) TableName() string {
	return "ticket_comments"
}
