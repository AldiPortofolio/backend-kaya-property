package auth

import (
	"kaya-backend/models"

	"kaya-backend/api/v1/services/auth"
	"kaya-backend/api/v1/services/customer"

	"github.com/gin-gonic/gin"
)

type (
	authontroller struct {
		Gen             *models.GeneralModel
		AuthService     auth.ServiceAuth
		CustomerService customer.CustomerService
	}

	AuthController interface {
		Auth(ctx *gin.Context)
	}
)

func InitiateAuthInterface(gen *models.GeneralModel) *authontroller {
	return &authontroller{
		Gen:             gen,
		AuthService:     *auth.InitiateServiceAuth(gen),
		CustomerService: *customer.InitiateCustomerInterface(gen),
	}
}
