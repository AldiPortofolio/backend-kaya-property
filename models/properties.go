package models

import (
	"time"

	"gorm.io/gorm"
)

type Properties struct {
	gorm.Model
	Slug                     string           `gorm:"column:slug" json:"slug"`
	Name                     string           `gorm:"column:name" json:"name"`
	Description              string           `gorm:"column:description" json:"description"`
	Price                    float64          `gorm:"column:price" json:"price"`
	Lot                      int              `gorm:"column:lot" json:"lot"`
	LotAvailable             int              `gorm:"column:lot_available" json:"lot_available"`
	PricePerLot              float64          `gorm:"column:price_per_lot" json:"price_per_lot"`
	EstimatedSellingPriceMin float64          `gorm:"column:estimated_selling_price_min" json:"estimated_selling_price_min"`
	EstimatedSellingPriceMax float64          `gorm:"column:estimated_selling_price_max" json:"estimated_selling_price_max"`
	Prospektus               string           `gorm:"column:prospektus" json:"prospektus"`
	SoldPrice                float64          `gorm:"column:sold_price" json:"sold_price"`
	SoldDate                 time.Time        `gorm:"column:sold_date" json:"sold_date"`
	CityID                   int              `gorm:"column:city_id" json:"city_id"`
	Presentase               int              `gorm:"column:presentase" json:"presentase"`
	IsSold                   bool             `gorm:"column:is_sold" json:"is_sold"`
	City                     City             `gorm:"references:CityID" json:"city"`
	PropertyFee              []PropertyFees     `gorm:"foreignKey:PropertyID;references:id" json:"property_fee"`
	PropertyPhotos           []PropertyPhotos `gorm:"foreignKey:PropertyID" json:"property_photos"`
}

type DetailProperties struct {
	gorm.Model
	Slug                     string           `gorm:"column:slug" json:"slug"`
	Name                     string           `gorm:"column:name" json:"name"`
	Description              string           `gorm:"column:description" json:"description"`
	Price                    float64          `gorm:"column:price" json:"price"`
	Lot                      int              `gorm:"column:lot" json:"lot"`
	LotAvailable             int              `gorm:"column:lot_available" json:"lot_available"`
	PricePerLot              float64          `gorm:"column:price_per_lot" json:"price_per_lot"`
	EstimatedSellingPriceMin float64          `gorm:"column:estimated_selling_price_min" json:"estimated_selling_price_min"`
	EstimatedSellingPriceMax float64          `gorm:"column:estimated_selling_price_max" json:"estimated_selling_price_max"`
	Prospektus               string           `gorm:"column:prospektus" json:"prospektus"`
	SoldPrice                float64          `gorm:"column:sold_price" json:"sold_price"`
	SoldDate                 time.Time        `gorm:"column:sold_date" json:"sold_date"`
	CityID                   int              `gorm:"column:city_id" json:"city_id"`
	LotSold                  float64          `json:"lot_sold"`
	IsSold                   bool             `gorm:"column:is_sold" json:"is_sold"`
	Presentase               int              `gorm:"column:presentase" json:"presentase"`
	City                     City             `gorm:"references:CityID" json:"city"`
	PropertyFee              []PropertyFees     `gorm:"foreignKey:PropertyID;references:id" json:"property_fee"`
	PropertyPhotos           []PropertyPhotos `gorm:"foreignKey:PropertyID" json:"property_photos"`
}

func (b *Properties) TableName() string {
	return "properties"
}

func (b *DetailProperties) TableName() string {
	return "properties"
}
