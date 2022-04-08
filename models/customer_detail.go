package models

type CustomerDetails struct {
	CustomerID     int    `gorm:"column:customer_id" json:"customer_id"`
	CityID         int    `gorm:"column:city_id" json:"city_id"`
	IdentityNumber string `gorm:"column:identity_number" json:"identity_number"`
	Address        string `gorm:"column:address" json:"address"`
	Selfie         string `gorm:"column:selfie" json:"selfie"`
	IdentityPhoto  string `gorm:"column:identity_photo" json:"identity_photo"`
	DatabaseModel
}

func (t *CustomerDetails) TableName() string {
	return "customer_details"
}
