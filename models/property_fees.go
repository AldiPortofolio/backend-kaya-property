package models

import (
	"gorm.io/gorm"
)

type PropertyFees struct {
	gorm.Model
	PropertyID int     `gorm:"column:property_id" json:"property_id"`
	FeeType    string  `gorm:"column:fee_type" json:"fee_type"`
	Amount     float64 `gorm:"column:amount" json:"amount"`
}

func (t *PropertyFees) TableName() string {
	return "property_fees"
}
