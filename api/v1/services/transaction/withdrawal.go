package transaction

import (
	"encoding/json"
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/helper"
)

type RekeningTujuan struct {
	AccountId  uint   `json:"account_id"`
	AnRekening string `json:"an_rekening"`
	NoRekening string `json:"no_rekening"`
	Bank       string `json:"bank"`
}

func (svc *TransactionService) Withdrawal(customerID uint, req request.Withdrawal, res *models.Response) {
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

	noTransction, _ := svc.TransactionRepo.GetRunNum("inv")
	token, err := helper.GenerateToken(noTransction)

	rekeningTujuan := RekeningTujuan{
		AccountId:  customer.CustomerAccounts.ID,
		AnRekening: customer.CustomerAccounts.Name,
		NoRekening: customer.CustomerAccounts.AccountNumber,
		Bank:       customer.CustomerAccounts.Bank.Name,
	}

	additionalData, err := json.Marshal(rekeningTujuan)
	if err != nil {
		fmt.Println("error", err)
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "Silahkan hubungi administrator"
		return
	}

	paymentMethod, _ := svc.TransactionRepo.FindPaymentMethodByName("BCA")

	transaction := models.Transactions{
		CustomerID:      int(customerID),
		SubTotal:        req.Amount,
		GrandTotal:      req.Amount,
		Fee:             0,
		StatusId:        constants.PENDING,
		PaymentMethodId: paymentMethod.ID,
		NoTransaction:   noTransction,
		TransactionType: constants.WITHDRAWAL,
		Token:           token,
		AdditionalData:  fmt.Sprintf("%s", additionalData),
	}

	_, err = svc.TransactionRepo.Save(transaction)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "Silahkan hubungi administrator"
		return
	}

	res.Data = token
	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "Withdrawal berhasil"
}
