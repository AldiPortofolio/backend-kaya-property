package titipjual

import (
	"kaya-backend/models"
	"kaya-backend/utils/constants"
)

func (svc *TitipJualService) Save(req models.TitipJual, res *models.Response) {
	//Save Titip Jual
	_, err := svc.titipJualRepo.Save(req)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
}
