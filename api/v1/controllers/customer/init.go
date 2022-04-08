package customer

import (
	"kaya-backend/models"

	"kaya-backend/api/v1/services/customer"

	"github.com/gin-gonic/gin"
)

type (
	customerController struct {
		Gen             *models.GeneralModel
		CustomerService customer.CustomerService
	}

	CustomerController interface {
		Customer(ctx *gin.Context)
	}
)

func InitiateCustomerInterface(gen *models.GeneralModel) *customerController {
	return &customerController{
		Gen:             gen,
		CustomerService: *customer.InitiateCustomerInterface(gen),
	}
}
