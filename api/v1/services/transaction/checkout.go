package transaction

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/repository"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/helper"
	"time"
)

func (svc *TransactionService) Checkout(req request.Payment, res *models.Response) {
	//Get Data Customer
	customer, err := svc.CustomerRepo.FindByID(uint64(req.CustomerID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	//Get Data Property
	property, err := svc.PropertyRepo.Find(req.PropertyID)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	if property.LotAvailable <= 0 {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "Lot sudah habis"
		return
	}

	//Get CustomerLot
	customerPropertyLotExist, err := svc.CustomerPropertyLotRepo.FindByPropertyID(int(property.ID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	totalLot := 0
	for _, obj := range customerPropertyLotExist {
		totalLot += obj.Lot
	}

	//Update Property Lot Available
	property.LotAvailable = property.LotAvailable - req.Lot
	devide := float32(0)
	if totalLot > 0 {
		devide = (float32(totalLot) + float32(req.Lot)) / float32(property.Lot)
	} else {
		devide = float32(req.Lot) / float32(property.Lot)
	}
	presentase := devide * 100
	property.Presentase = int(presentase)

	//Set Data for Transaction Detail
	transactionDetail := models.TransactionDetails{
		PropertyID: int(property.ID),
		Lot:        req.Lot,
		Price:      property.PricePerLot,
	}

	status := constants.PENDING
	paymentMethod, _ := svc.TransactionRepo.FindPaymentMethodByName(req.PaymentMethode)
	//Set Data for Transaction
	subTotal := float64(property.PricePerLot * float64(req.Lot))
	grandTotal := float64(property.PricePerLot * float64(req.Lot))

	fmt.Println("balance", customer.BalanceAmount >= grandTotal)

	randomNumber := 0

	if paymentMethod.Code == "saldo" && customer.BalanceAmount < grandTotal {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "Saldo tidak cukup"
		return
	} else if paymentMethod.Code == "saldo" && customer.BalanceAmount >= grandTotal {
		status = constants.SUCCESS
	} else if paymentMethod.Code == "bca" {
		randomNumber = svc.TransactionRepo.GetRandomNumber()
		grandTotal = grandTotal + float64(randomNumber)
	}

	noTransction, _ := svc.TransactionRepo.GetRunNum("inv")

	token, err := helper.GenerateToken(noTransction)

	if err != nil {
		res.Meta.Code = constants.INTERNAL_SERVER_ERROR_CODE
		res.Meta.Status = false
		res.Meta.Message = "Error, Silahkan hubungi administrator"
		res.Meta.ErrorMessage = err.Error()
		return
	}

	transaction := models.Transactions{
		CustomerID:        int(customer.ID),
		SubTotal:          subTotal,
		GrandTotal:        grandTotal,
		Fee:               0,
		TransactionDetail: transactionDetail,
		StatusId:          status,
		PaymentMethodId:   paymentMethod.ID,
		NoTransaction:     noTransction,
		TransactionType:   constants.PEMBELIAN_LOT,
		Token:             token,
		RandomNumber:      randomNumber,
	}

	// Start of Transaction
	trxHandle := repository.Dbcon.Begin()
	// Save Transaction
	resTransaction, errTransaction := svc.TransactionRepo.WithTrx(trxHandle).Save(transaction)
	if errTransaction != nil {
		trxHandle.Rollback()
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "Error, silahkan hubungi administrator"
		return
	}

	if paymentMethod.Code == "saldo" {
		//Set Data for Balance Transaction
		balanceTransaction := models.BalanceTransaction{
			CustomerID:      int(customer.ID),
			TransactionType: constants.PEMBELIAN_LOT,
			Description:     property.Name,
			Amount:          grandTotal,
			Status:          constants.STATUS_DONE,
			Date:            time.Now(),
			Note:            "",
			Lot:             req.Lot,
		}
		// Save Balance transaction
		_, errBalanceTrasaction := svc.BalanceTransactionRepo.WithTrx(trxHandle).Save(balanceTransaction)
		if errBalanceTrasaction != nil {
			trxHandle.Rollback()
			res.Meta.Code = constants.BAD_REQUEST_CODE
			res.Meta.Status = false
			res.Meta.Message = "Error, silahkan hubungi administrator"
			return
		}

		existingCustomerPropertyLot, err := svc.CustomerPropertyLotRepo.FindByCustomerIDAndPropertyId(int(customer.ID), req.PropertyID)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Set Data Customer Property Lot
			customerPropertyLot := models.CustomerPropertyLots{
				CustomerID:          int(customer.ID),
				PropertyID:          int(property.ID),
				Lot:                 req.Lot,
				TransactionDetailID: resTransaction.TransactionDetail.ID,
			}

			// Save Customer Property Lot
			_, errCustomerPropertyLot := svc.CustomerPropertyLotRepo.WithTrx(trxHandle).Save(customerPropertyLot)
			if errCustomerPropertyLot != nil {
				trxHandle.Rollback()
				res.Meta.Code = constants.BAD_REQUEST_CODE
				res.Meta.Status = false
				res.Meta.Message = "Error, silahkan hubungi administrator"
				return
			}
		} else {
			existingCustomerPropertyLot.Lot = req.Lot + existingCustomerPropertyLot.Lot
			// Save Customer Property Lot
			errCustomerPropertyLot := svc.CustomerPropertyLotRepo.WithTrx(trxHandle).UpdateCustomerPropertyLot(existingCustomerPropertyLot)
			if errCustomerPropertyLot != nil {
				trxHandle.Rollback()
				res.Meta.Code = constants.BAD_REQUEST_CODE
				res.Meta.Status = false
				res.Meta.Message = "Error, silahkan hubungi administrator"
				return
			}
		}

		customer.BalanceAmount = customer.BalanceAmount - grandTotal
		_, errUpdateCustomer := svc.CustomerRepo.WithTrx(trxHandle).UpdateBalance(customer)
		if errUpdateCustomer != nil {
			trxHandle.Rollback()
			res.Meta.Code = constants.BAD_REQUEST_CODE
			res.Meta.Status = false
			res.Meta.Message = "Error, silahkan hubungi administrator"
			return
		}

	} else {

	}

	//// Update Property
	_, errProperty := svc.PropertyRepo.WithTrx(trxHandle).Save(property) //(transaction, customerPropertyLot, property)
	if errProperty != nil {
		trxHandle.Rollback()
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "Error, silahkan hubungi administrator"
		return
	}

	//Check to Mutation Bank BCA
	//checkMutationBank()

	// End of transaction
	trxHandle.Commit()

	res.Data = token
	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
}

func checkMutationBank() {
	fmt.Println("Check Mutation Bank BCA")
}
