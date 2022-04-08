package transaction

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/helper"
)

func (svc *TransactionService) Detail(noOrder string, res *models.Response) {
	fmt.Println(">>> Transaction Detail - Controller <<<")

	no, err := helper.DecryptToken(noOrder)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "No transaksi tidak ditemukan"
		res.Meta.ErrorMessage = err.Error()
		return
	}

	data, err := svc.TransactionRepo.FindByNoOrder(no)

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
