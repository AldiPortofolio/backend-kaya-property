package dashboard

import (
	"kaya-backend/models"
	"kaya-backend/utils/constants"
)

func (svc *DashboardService) CancelSell(ID int, res *models.Response) {
	data, err := svc.CustomerPropertySecondaryRepo.FindByID(ID)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	data.Status = constants.SECONDARY_STATUS_BATAL

	errRes := svc.CustomerPropertySecondaryRepo.Delete(data)
	if errRes != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
}
