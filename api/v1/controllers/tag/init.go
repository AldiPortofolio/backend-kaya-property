package tag

import (
	"kaya-backend/models"

	tag "kaya-backend/api/v1/services/tag"

	"github.com/gin-gonic/gin"
)

type (
	tagController struct {
		Gen        *models.GeneralModel
		TagService tag.TagService
	}

	TagController interface {
		Tag(ctx *gin.Context)
	}
)

func InitiateTagInterface(gen *models.GeneralModel) *tagController {
	return &tagController{
		Gen:        gen,
		TagService: *tag.InitiateTagService(gen),
	}
}
