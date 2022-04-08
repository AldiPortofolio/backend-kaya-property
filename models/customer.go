package models

type Customers struct {
	Name              string           `gorm:"column:name" json:"name"`
	Email             string           `gorm:"column:email" json:"email"`
	Password          string           `gorm:"column:password" json:"password"`
	NoHp              string           `gorm:"column:no_handphone" json:"no_handphone"`
	VerifyCode        string           `gorm:"column:verify_code" json:"verify_code"`
	BalanceAmount     float64          `gorm:"column:balance_amount" json:"balance_amount"`
	IsActive          bool             `gorm:"column:is_active" json:"is_active"`
	CustomerDetails   CustomerDetails  `gorm:"foreignKey:CustomerID;references:id" json:"customer_details"`
	CustomerAccounts  CustomerAccounts `gorm:"foreignKey:CustomerID;references:id" json:"customer_accounts"`
	MembershipLevelID uint             `gorm:"column:membership_level_id" json:"membership_level_id"`
	DatabaseModel
}

type Me struct {
	Name              string           `gorm:"column:name" json:"name"`
	Email             string           `gorm:"column:email" json:"email"`
	NoHp              string           `gorm:"column:no_handphone" json:"no_handphone"`
	BalanceAmount     float64          `gorm:"column:balance_amount" json:"balance_amount"`
	CustomerDetails   CustomerDetails  `gorm:"foreignKey:CustomerID;references:id" json:"customer_details"`
	CustomerAccounts  CustomerAccounts `gorm:"foreignKey:CustomerID;references:id" json:"customer_accounts"`
	MembershipLevelID uint             `gorm:"column:membership_level_id" json:"membership_level_id"`
	DatabaseModel
}

func (t *Customers) TableName() string {
	return "customers"
}

func (t *Me) TableName() string {
	return "customers"
}

type Guest struct {
	Email string `gorm:"column:email" json:"email"`
	DatabaseModel
}

func (t *Guest) TableName() string {
	return "guest_email"
}
