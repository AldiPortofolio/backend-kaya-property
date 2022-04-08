package repository

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/helper"
	"time"

	"github.com/jinzhu/gorm"
)

func NewAuthRepository(gen *models.GeneralModel, db *gorm.DB) *authRepository {
	return &authRepository{
		General: gen,
		DB:      db,
	}
}

type (
	AuthRepository interface {
		Login(req request.Login) (res models.Customers, result bool, err error)
		ByEmail(req request.ByEmail) (res models.Customers, err error)
	}
	authRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo authRepository) Login(req request.Login) (res models.Customers, result bool, err error) {
	fmt.Println(">>> Database - Login <<<")
	defer timeTrack(time.Now(), "Login")

	err = repo.DB.Where("email = ? ", req.Email).First(&res).Error
	if err != nil {
		return res, true, err
	}

	correct := helper.CheckPasswordHash(req.Password, res.Password)
	fmt.Println(correct)

	if !correct {
		return res, false, nil
	}

	return res, true, nil
}

func (repo authRepository) ByEmail(req request.ByEmail) (res models.Customers, err error) {
	fmt.Println(">>> Database - ByEmail <<<")
	defer timeTrack(time.Now(), "ByEmail")

	err = repo.DB.Where("email = ? ", req.Email).First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
