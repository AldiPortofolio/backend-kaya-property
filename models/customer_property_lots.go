package models

type CustomerPropertyLots struct {
	ID                  int        `gorm:"column:id" json:"id"`
	CustomerID          int        `gorm:"column:customer_id" json:"customer_id"`
	PropertyID          int        `gorm:"column:property_id" json:"property_id"`
	TransactionDetailID int        `gorm:"column:transaction_detail_id" json:"transaction_detail_id"`
	Lot                 int        `gorm:"column:lot" json:"lot"`
	Property            Properties `gorm:"references:PropertyID" json:"property"`
}

func (t *CustomerPropertyLots) TableName() string {
	return "customer_property_lots"
}
