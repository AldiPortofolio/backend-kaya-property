package customer

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/helper"
)

func (svc *CustomerService) Upload(customerID uint, req request.CustomerUpload, res *models.Response) {
	customer, err := svc.CustomerRepo.FindByID(uint64(customerID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	verifyCode := helper.GenerateRandomCode(8)
	customer.VerifyCode = verifyCode

	customer.CustomerDetails.Selfie = req.UploadSelfie
	customer.CustomerDetails.IdentityPhoto = req.UploadKtp
	result, err := svc.CustomerRepo.Save(customer)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	err = helper.SendVerification(customer.NoHp, customer.VerifyCode)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "Upload file berhasil"
	res.Data = result
}
