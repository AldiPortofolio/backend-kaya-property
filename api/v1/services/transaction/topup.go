package transaction

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/helper"
)

func (svc *TransactionService) Topup(customerID uint, req request.Topup, res *models.Response) {
	_, err := svc.CustomerRepo.FindByID(uint64(customerID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	paymentMethod, _ := svc.TransactionRepo.FindPaymentMethodByName("BCA")

	noTransction, _ := svc.TransactionRepo.GetRunNum("inv")
	token, err := helper.GenerateToken(noTransction)
	randomNumber := 0
	randomNumber = svc.TransactionRepo.GetRandomNumber()
	grandTotal := req.Amount + float64(randomNumber)

	transaction := models.Transactions{
		CustomerID:      int(customerID),
		SubTotal:        req.Amount,
		GrandTotal:      grandTotal,
		Fee:             0,
		PaymentMethodId: paymentMethod.ID,
		StatusId:        constants.PENDING,
		NoTransaction:   noTransction,
		TransactionType: constants.TOPUP,
		Token:           token,
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
	res.Meta.Message = "Topup berhasil"
}
