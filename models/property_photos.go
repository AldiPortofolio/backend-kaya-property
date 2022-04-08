package models

import (
	"gorm.io/gorm"
)

type PropertyPhotos struct {
	gorm.Model
	PropertyID int    `gorm:"column:property_id" json:"property_id"`
	Type       string `gorm:"column:type" json:"type"`
	Photo      string `gorm:"column:photo" json:"photo"`
}

func (t *PropertyPhotos) TableName() string {
	return "property_photos"
}
