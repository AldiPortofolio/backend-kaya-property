package properties

import (
	"kaya-backend/models"

	customerpropertysecondary "kaya-backend/api/v1/services/customer-property-secondary"
	"kaya-backend/api/v1/services/properties"

	"github.com/gin-gonic/gin"
)

type (
	propertiesController struct {
		Gen                              *models.GeneralModel
		PropertiesService                properties.PropertyService
		CustomerPropertySecondaryService customerpropertysecondary.CustomerPropertySecondaryService
	}

	PropertiesController interface {
		Property(ctx *gin.Context)
	}
)

func InitiatePropertiesInterface(gen *models.GeneralModel) *propertiesController {
	return &propertiesController{
		Gen:                              gen,
		PropertiesService:                *properties.InitiatePropretyInterface(gen),
		CustomerPropertySecondaryService: *customerpropertysecondary.InitiateCustomerPropertySecondaryInterface(gen),
	}
}
