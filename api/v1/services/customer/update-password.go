package customer

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/helper"
)

func (svc *CustomerService) UpdatePassword(customerID uint, req request.CustomerPassword, res *models.Response) {
	customer, err := svc.CustomerRepo.FindByID(uint64(customerID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	correct := helper.CheckPasswordHash(req.OldPassword, customer.Password)
	fmt.Println("correct", correct)
	if !correct {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "old password not match"
		return
	}

	hashedPassword, errPass := helper.HashPassword(req.NewPassword)
	if errPass != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = errPass.Error()
		return
	}

	customer.Password = hashedPassword
	_, errRes := svc.CustomerRepo.Save(customer)
	if errRes != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "Password berhasil diupdate"
}
