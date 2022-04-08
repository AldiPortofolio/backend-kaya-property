package models

type CustomerAccounts struct {
	CustomerID    int    `gorm:"column:customer_id" json:"customer_id"`
	BankID        int    `gorm:"column:bank_id" json:"bank_id"`
	Name          string `gorm:"column:name" json:"name"`
	AccountNumber string `gorm:"column:account_number" json:"account_number"`
	Bank          Bank   `gorm:"references:BankID;" json:"bank"`
	DatabaseModel
}

func (t *CustomerAccounts) TableName() string {
	return "customer_accounts"
}
