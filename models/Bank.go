package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
}

func (b *Bank) TableName() string {
	return "banks"
}
