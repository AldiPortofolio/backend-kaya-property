package repository

import (
	"fmt"
	"kaya-backend/models"
	"time"

	"github.com/jinzhu/gorm"
)

func NewMembershipLevelRepository(gen *models.GeneralModel, db *gorm.DB) *membershipLevelRepository {
	return &membershipLevelRepository{
		General: gen,
		DB:      db,
	}
}

type (
	MembershipLevelRepository interface {
		GetLevelDefault() (models.MembershipLevel, error)
	}
	membershipLevelRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo membershipLevelRepository) GetLevelDefault() (models.MembershipLevel, error) {
	fmt.Println(">>> Database - Membership Level <<<")
	defer timeTrack(time.Now(), "Membership Level")

	res := models.MembershipLevel{}

	err := repo.DB.Where("is_default = ?", true).Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, err
}
