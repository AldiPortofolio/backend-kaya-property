package models

type Status struct {
	ID   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Code string `gorm:"column:name" json:"code"`
}

func (s *Status) TableName() string {
	return "statuses"
}
