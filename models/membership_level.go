package models

import "gorm.io/gorm"

type MembershipLevel struct {
	gorm.Model
	Name                string  `gorm:"column:name" json:"name"`
	InvestUpTo          float64 `gorm:"invest_up_to" json:"investUpTo"`
	ManagementFeeType   string  `gorm:"management_fee_type" json:"ManagementFeeType"`
	ManagementFeeAmount float64 `gorm:"management_fee_amount" json:"ManagementFeeAmount"`
	FeeDisplay          float64 `gorm:"fee_display" json:"FeeDisplay"`
	FeeReal             float64 `gorm:"fee_real" json:"FeeReal"`
	IsDefault           bool    `gorm:"is_default" json:"is_default"`
}

func (t *MembershipLevel) TableName() string {
	return "membership_level"
}
