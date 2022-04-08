package models

type Province struct {
	ID   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (p *Province) TableName() string {
	return "provinces"
}
