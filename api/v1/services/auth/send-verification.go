package auth

import (
	"kaya-backend/models"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/helper"
)

func (svc *ServiceAuth) SendVerification(customerID uint, res *models.Response) {
	customer, err := svc.CustomerRepo.FindByID(uint64(customerID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	verifyCode := helper.GenerateRandomCode(8)
	customer.VerifyCode = verifyCode
	_, errRes := svc.CustomerRepo.Save(customer)
	if errRes != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	//err = helper.SendVerification(customer.NoHp, customer.VerifyCode)
	//if err != nil {
	//	res.Meta.Code = constants.BAD_REQUEST_CODE
	//	res.Meta.Status = false
	//	res.Meta.Message = err.Error()
	//	return
	//}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "Verification telah terkirim"
}
