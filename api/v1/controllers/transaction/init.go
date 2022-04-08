package transaction

import (
	"kaya-backend/models"

	transaction "kaya-backend/api/v1/services/transaction"

	"github.com/gin-gonic/gin"
)

type (
	transactionController struct {
		Gen                *models.GeneralModel
		TransactionService transaction.TransactionService
	}

	TransactionController interface {
		Transaction(ctx *gin.Context)
	}
)

func InitiateTransactionInterface(gen *models.GeneralModel) *transactionController {
	return &transactionController{
		Gen:                gen,
		TransactionService: *transaction.InitiateTransactionInterface(gen),
	}
}
