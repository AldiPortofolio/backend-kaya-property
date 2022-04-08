package models

import "time"

type Transactions struct {
	ID                int                `gorm:"column:id" json:"id"`
	CustomerID        int                `gorm:"column:customer_id" json:"customer_id"`
	SubTotal          float64            `gorm:"column:sub_total" json:"sub_total"`
	Fee               float64            `gorm:"column:fee" json:"fee"`
	GrandTotal        float64            `gorm:"column:grand_total" json:"grand_total"`
	TransactionDetail TransactionDetails `gorm:"foreignKey:TransactionID;references:id" json:"transaction_detail"`
	PaymentMethod     PaymentMethod      `gorm:"foreignKey:PaymentMethodId;references:id" json:"payment_method"`
	PaymentMethodId   int                `gorm:"column:payment_method_id" json:"payment_method_id"`
	StatusId          int                `gorm:"column:status_id" json:"status_id"`
	Status            Status             `gorm:"foreignKey:StatusId;references:id" json:"status"`
	NoTransaction     string             `gorm:"column:no_transaction" json:"no_transaction"`
	TransactionType   string             `gorm:"column:transaction_type" json:"transaction_type"`
	Token             string             `gorm:"column:token" json:"token"`
	AdditionalData    string             `gorm:"column:additional_data" json:"additional_data"`
	RandomNumber      int                `gorm:"column:random_number" json:"random_number"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (t *Transactions) TableName() string {
	return "transactions"
}

type TransactionDetails struct {
	ID            int        `gorm:"column:id" json:"id"`
	TransactionID int        `gorm:"column:transaction_id" json:"transaction_id"`
	PropertyID    int        `gorm:"column:property_id" json:"property_id"`
	Lot           int        `gorm:"column:lot" json:"lot"`
	Price         float64    `gorm:"column:price" json:"price"`
	Property      Properties `gorm:"foreignKey:PropertyId;references:id" json:"property"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (t *TransactionDetails) TableName() string {
	return "transaction_details"
}

type RunNums struct {
	ID             int    `gorm:"column:id" json:"id"`
	Prefix         string `gorm:"column:prefix" json:"prefix"`
	CodeFormat     string `gorm:"column:code_format" json:"code_format"`
	CodeLen        string `gorm:"column:code_len" json:"code_len"`
	LastNum        string `gorm:"column:last_num" json:"last_num"`
	Desc           string `gorm:"column:desc" json:"desc"`
	NumCode        string `gorm:"column:num_code" json:"num_code"`
	LastCodeFormat string `gorm:"column:last_code_format" json:"last_code_format"`
}

func (r *RunNums) TableName() string {
	return "tb_run_nums"
}

type RandomNumber struct {
	ID     int `gorm:"column:id" json:"id"`
	Number int `gorm:"column:prefix" json:"prefix"`
}

func (r *RandomNumber) TableName() string {
	return "random_number"
}
