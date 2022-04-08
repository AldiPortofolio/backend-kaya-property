package customer

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
)

func (svc *CustomerService) UpdateCustomerAccount(customerID uint, req request.CustomerAccount, res *models.Response) {
	customer, err := svc.CustomerRepo.FindByID(uint64(customerID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	customer.CustomerAccounts.BankID = req.BankID
	customer.CustomerAccounts.Name = req.Name
	customer.CustomerAccounts.AccountNumber = req.AccountNumber

	_, errRes := svc.CustomerRepo.Save(customer)
	if errRes != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "Data berhasil diupdate"
}
