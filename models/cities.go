package models

type City struct {
	ID         int      `gorm:"column:id" json:"id"`
	Name       string   `gorm:"column:name" json:"name"`
	ProvinceID int      `gorm:"column:province_id" json:"province_id"`
	Province   Province `gorm:"references:ProvinceID;" json:"province"`
}

func (c *City) TableName() string {
	return "cities"
}
