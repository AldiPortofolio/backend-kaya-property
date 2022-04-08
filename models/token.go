package models

type Tokens struct {
	ID         int    `gorm:"column:id" json:"id"`
	Token      string `gorm:"column:token" json:"token"`
	CustomerID int    `gorm:"column:customer_id" json:"customer_id"`
	IsActive   bool   `gorm:"column:is_active" json:"is_active"`
}

func (t *Tokens) TableName() string {
	return "tokens"
}
