package blog

import (
	"kaya-backend/models"

	blog "kaya-backend/api/v1/services/blog"

	"github.com/gin-gonic/gin"
)

type (
	blogController struct {
		Gen        *models.GeneralModel
		BlogService blog.BlogService
	}

	BlogController interface {
		Blog(ctx *gin.Context)
	}
)

func InitiateBlogInterface(gen *models.GeneralModel) *blogController {
	return &blogController{
		Gen:        gen,
		BlogService: *blog.InitiateBlogService(gen),
	}
}
