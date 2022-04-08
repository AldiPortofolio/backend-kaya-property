package auth

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/repository"

	logger "kaya-backend/library/logger/v2"
)

// ServiceAuth ..
type ServiceAuth struct {
	General      *models.GeneralModel
	OttoLog      logger.KayalogInterface
	AuthRepo     repository.AuthRepository
	CustomerRepo repository.CustomerRepository
	TokenRepo    repository.TokenRepository
	Database     repository.DbPostgres
}

// ServicLoginInterface ..
type ServiceAuthInterface interface {
	Login(request.Login, *models.Response)
	RefreshToken(request.RefreshToken, *models.Response)
}

// InitiateServicActivitasSalesmen ..
func InitiateServiceAuth(gen *models.GeneralModel) *ServiceAuth {
	return &ServiceAuth{
		General:      gen,
		AuthRepo:     repository.NewAuthRepository(gen, repository.Dbcon),
		CustomerRepo: repository.NewCustomerRepository(gen, repository.Dbcon),
		TokenRepo:    repository.NewTokenRepository(gen, repository.Dbcon),
	}
}
