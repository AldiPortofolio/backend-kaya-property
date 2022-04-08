package customer

import (
	"kaya-backend/models"
	"kaya-backend/utils/constants"
)

func (svc *CustomerService) Me(customerID uint, res *models.Response) {
	customer, err := svc.CustomerRepo.FindMe(uint64(customerID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
	res.Data = customer
}
