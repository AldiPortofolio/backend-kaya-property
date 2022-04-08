package customerpropertysecondary

import (
	"kaya-backend/models"
	"kaya-backend/utils/constants"
)

func (svc *CustomerPropertySecondaryService) DetailSecondary(slug string, res *models.Response) {
	data, err := svc.PropertyRepo.FindBySlug(slug)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
	res.Data = data
}

func (svc *CustomerPropertySecondaryService) DetailSecondaryByID(ID int, res *models.Response) {
	data, err := svc.CustomerPropertySecondaryRepo.FindByID(ID)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
	res.Data = data
}
