package models

type CustomerPropertySecondaries struct {
	ID                    int                  `gorm:"column:id" json:"id"`
	CustomerID            int                  `gorm:"column:customer_id" json:"customer_id"`
	PropertyID            int                  `gorm:"column:property_id" json:"property_id"`
	Lot                   int                  `gorm:"column:lot" json:"lot"`
	PricePerLot           float64              `gorm:"column:price_per_lot" json:"price_per_lot"`
	Status                string               `gorm:"column:status" json:"status"`
	Property              Properties           `gorm:"references:PropertyID" json:"property"`
	Customer              Customers            `gorm:"references:CustomerID" json:"customer"`
	CustomerPropertyLotID int                  `gorm:"column:customer_property_lot_id" json:"customer_property_lot_id"`
	CustomerPropertyLot   CustomerPropertyLots `gorm:"references:CustomerPropertyLotID" json:"customer_property_lot"`
}

type ResSumCustomerPropertySecondaries struct {
	TotalLot   int     `json:"totalLot"`
	TotalPrice float64 `json:"totalPrice"`
}

func (t *CustomerPropertySecondaries) TableName() string {
	return "customer_property_secondaries"
}
