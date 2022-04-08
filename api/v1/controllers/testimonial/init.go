package testimonial

import (
	"kaya-backend/models"

	testimonials "kaya-backend/api/v1/services/testimonial"

	"github.com/gin-gonic/gin"
)

type (
testimonialController struct {
		Gen              *models.GeneralModel
		testimonialService testimonials.Testimonialervice
	}

	TestimonialController interface {
		Testimonials(ctx *gin.Context)
	}
)

func InitiateTestimonialInterface(gen *models.GeneralModel) *testimonialController {
	return &testimonialController{
		Gen:              gen,
		testimonialService: *testimonials.InitiateTestimonialInterface(gen),
	}
}
