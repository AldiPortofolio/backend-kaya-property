package models

type Tag struct {
	ID   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (t *Tag) TableName() string {
	return "tags"
}
