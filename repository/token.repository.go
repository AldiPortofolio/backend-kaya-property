package repository

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"time"

	"github.com/jinzhu/gorm"
)

// NewTokenRepository ..
func NewTokenRepository(gen *models.GeneralModel, db *gorm.DB) *tokenRepository {
	return &tokenRepository{
		General: gen,
		DB:      db,
	}
}

// TransactionRepository ..
type (
	TokenRepository interface {
		Save(models.Tokens) (models.Tokens, error)
		FindByFilter(req request.ResetPassword) (res models.Tokens, err error)
	}
	tokenRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo tokenRepository) Save(req models.Tokens) (models.Tokens, error) {
	fmt.Println(">>> tokenRepository - Save <<<")
	defer timeTrack(time.Now(), "Save")

	err := repo.DB.Save(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

func (repo tokenRepository) FindByFilter(req request.ResetPassword) (res models.Tokens, err error) {
	fmt.Println(">>> tokenRepository - FindByFilter <<<")
	defer timeTrack(time.Now(), "FindByFilter")

	db := repo.DB
	db = db.Where("token = ? and customer_id = ? and is_active=true", req.Token, req.CustomerID).Order("created_at DESC")

	err = db.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
