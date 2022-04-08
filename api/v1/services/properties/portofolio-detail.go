package properties

import (
	"kaya-backend/models"
	"kaya-backend/utils/constants"
)

func (svc *PropertyService) PortofolioDetail(ID int, res *models.Response) {
	data, err := svc.customerPropertyLotRepo.FindByID(ID)
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
