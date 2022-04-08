package models

type PaymentMethod struct {
	ID             int    `gorm:"column:id" json:"id"`
	Name           string `gorm:"column:name" json:"name"`
	Code           string `gorm:"column:code" json:"code"`
	AdditionalData string `gorm:"column:additional_data" json:"additional_data"`
}

func (p *PaymentMethod) TableName() string {
	return "payment_method"
}
