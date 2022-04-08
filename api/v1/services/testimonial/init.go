package testimonial

import (
	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"
	"kaya-backend/repository"
)

// Testimonialervice ..
type Testimonialervice struct {
	General       *models.GeneralModel
	OttoLog       logger.KayalogInterface
	testimonialRepo repository.TestimonialsRepository
	Database      repository.DbPostgres
}

// InitiateTestimonialInterface ..
func InitiateTestimonialInterface(gen *models.GeneralModel) *Testimonialervice {
	return &Testimonialervice{
		General:       gen,
		testimonialRepo: repository.NewTestimonialsRepository(gen, repository.Dbcon),
	}
}
