package controllers

import (
	"kaya-backend/models"
	db "kaya-backend/repository"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ProvinceOptions(ctx *gin.Context) {
	provinces := []models.Province{}

	query := strings.ToLower(ctx.Query("name"))

	if dbRes := db.GetDbCon().Unscoped().Where("LOWER(name) like ?", "%"+query+"%").Find(&provinces); dbRes.Error != nil {
		log.Println("Get Province error : ", dbRes.Error)
		ctx.JSON(http.StatusBadRequest, "Error")
	}
	ctx.JSON(http.StatusOK, provinces)
}

func CityOptions(ctx *gin.Context) {
	cities := []models.City{}
	query := strings.ToLower(ctx.Query("province_id"))

	if dbRes := db.GetDbCon().Unscoped().Where("province_id = ?", query).Find(&cities); dbRes.Error != nil {
		log.Println("Get Province error : ", dbRes.Error)
		ctx.JSON(http.StatusBadRequest, "Error")
	}
	ctx.JSON(http.StatusOK, cities)
}

func AllCityOptions(ctx *gin.Context) {
	cities := []models.City{}

	if dbRes := db.GetDbCon().Unscoped().Find(&cities); dbRes.Error != nil {
		log.Println("Get Province error : ", dbRes.Error)
		ctx.JSON(http.StatusBadRequest, "Error")
	}
	ctx.JSON(http.StatusOK, cities)
}

func BankOptions(ctx *gin.Context) {
	banks := []models.Bank{}
	query := strings.ToLower(ctx.Query("name"))
	if dbRes := db.GetDbCon().Unscoped().Where("LOWER(name) like ?", "%"+query+"%").Find(&banks); dbRes.Error != nil {
		log.Println("Get Province error : ", dbRes.Error)
		ctx.JSON(http.StatusBadRequest, "Error")

	}
	ctx.JSON(http.StatusOK, banks)
}
