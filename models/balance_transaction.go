package models

import "time"

type BalanceTransaction struct {
	CustomerID      int       `gorm:"column:customer_id" json:"customer_id"`
	TransactionType string    `gorm:"column:transaction_type" json:"transaction_type"`
	Description     string    `gorm:"column:description" json:"description"`
	Amount          float64   `gorm:"column:amount" json:"amount"`
	Status          string    `gorm:"column:status" json:"status"`
	Date            time.Time `gorm:"column:date" json:"date"`
	Note            string    `gorm:"column:note" json:"note"`
	FollowedUpBy    string    `gorm:"column:follow_up_by" json:"follow_up_by"`
	Lot    					int    `gorm:"column:lot" json:"lot"`
	DatabaseModel
}

func (t *BalanceTransaction) TableName() string {
	return "balance_transactions"
}
