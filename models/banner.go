package models

type Banner struct {
	ID           int    `gorm:"column:id" json:"id"`
	Content      string `gorm:"column:content" json:"content"`
	Photo        string `gorm:"column:photo" json:"photo"`
	CallToAction string `gorm:"column:call_to_action" json:"call_to_action"`
}

func (t *Banner) TableName() string {
	return "banners"
}
