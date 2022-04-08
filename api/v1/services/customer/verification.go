package customer

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
)

func (svc *CustomerService) Verify(customerID uint, req request.CustomerVerification, res *models.Response) {
	customer, err := svc.CustomerRepo.FindByID(uint64(customerID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	if req.VerifyCode != customer.VerifyCode {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "Verification code invalid"
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "Verification code valid"
}
