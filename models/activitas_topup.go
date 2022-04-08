package models

import "time"

type HistoryTopup struct {
	Transaction string    `gorm:"column:transaction" json:"transaction"`
	Description string    `gorm:"column:description" json:"description"`
	Amount      float64   `gorm:"column:amount" json:"amount"`
	Status      string    `gorm:"column:status" json:"status"`
	Date        time.Time `gorm:"column:date" json:"date"`
}

type HistoryTransaction struct {
	Transaction string    `gorm:"column:transaction" json:"transaction"`
	Property    string    `gorm:"column:description" json:"description"`
	Lot         int       `gorm:"column:lot" json:"lot"`
	Amount      float64   `gorm:"column:amount" json:"amount"`
	Date        time.Time `gorm:"column:date" json:"date"`
}
