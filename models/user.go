package models

type User struct {
	ID       int    `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Email    string `gorm:"column:email" json:"email"`
}

func (p *User) TableName() string {
	return "users"
}
