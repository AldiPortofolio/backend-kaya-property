package models

import "time"

type Testimonials struct {
	ID       int    `gorm:"column:id" json:"id"`
	FullName string    `gorm:"column:full_name" json:"full_name"`
	Description  string `gorm:"column:description" json:"description"`
	Photo string `gorm:"column:photo" json:"photo"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (t *Testimonials) TableName() string {
	return "testimonials"
}
