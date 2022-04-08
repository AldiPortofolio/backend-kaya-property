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

func (svc *TransactionService) CheckoutSecondary(req request.Payment, res *models.Response) {
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

	//Get Data Customer Property Secondary
	customerPropertySecondary, err := svc.CustomerPropertySecondaryRepo.FindByID(req.SecondaryID)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	if customerPropertySecondary.Status != constants.SECONDARY_STATUS_OPEN {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "Lot sudah terjual atau sudah dibatalkan"
		return
	}

	//Get Data Customer Property Lot
	customerPropertyLot, err := svc.CustomerPropertyLotRepo.FindByID(customerPropertySecondary.CustomerPropertyLotID)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	status := constants.PENDING
	paymentMethod, _ := svc.TransactionRepo.FindPaymentMethodByName(req.PaymentMethode)

	subTotal := float64(customerPropertySecondary.PricePerLot * float64(req.Lot))
	fee := float64(subTotal * 0.0075)
	grandTotal := float64((customerPropertySecondary.PricePerLot * float64(req.Lot)) + fee)
	noTransction, _ := svc.TransactionRepo.GetRunNum("inv")
	token, err := helper.GenerateToken(noTransction)

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

	//Set Data for Transaction Detail
	transactionDetail := models.TransactionDetails{
		PropertyID: int(property.ID),
		Lot:        req.Lot,
		Price:      customerPropertySecondary.PricePerLot,
	}

	//Set Data for Transaction

	transaction := models.Transactions{
		CustomerID:        int(customer.ID),
		SubTotal:          subTotal,
		GrandTotal:        grandTotal,
		Fee:               fee,
		TransactionDetail: transactionDetail,
		StatusId:          status,
		PaymentMethodId:   paymentMethod.ID,
		NoTransaction:     noTransction,
		TransactionType:   constants.PEMBELIAN_LOT,
		Token:             token,
	}

	////Set Data for Balance Transaction Pembeli
	//balanceTransactionPembeli := models.BalanceTransaction{
	//	CustomerID:      int(customer.ID),
	//	TransactionType: constants.PEMBELIAN_LOT,
	//	Description:     property.Name,
	//	Amount:          grandTotal,
	//	Status:          constants.STATUS_PENDING,
	//	Date:            time.Now(),
	//	Note:            "",
	//	Lot:             req.Lot,
	//}
	//
	////Set Data for Balance Transaction Penjual
	//balanceTransactionPenjual := models.BalanceTransaction{
	//	CustomerID:      customerPropertySecondary.CustomerID,
	//	TransactionType: constants.PENJUALAN_PROPERTY,
	//	Description:     property.Name,
	//	Amount:          grandTotal,
	//	Status:          constants.STATUS_PENDING,
	//	Date:            time.Now(),
	//	Note:            "",
	//	Lot:             req.Lot,
	//}

	// Start of Transaction
	trxHandle := repository.Dbcon.Begin()
	// Save Transaction
	resTransaction, errTransaction := svc.TransactionRepo.WithTrx(trxHandle).Save(transaction)
	if errTransaction != nil {
		trxHandle.Rollback()
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = errTransaction.Error()
		return
	}

	//// Save Balance transaction pembeli
	//_, errBalanceTrasactionPembeli := svc.BalanceTransactionRepo.WithTrx(trxHandle).Save(balanceTransactionPembeli)
	//if errBalanceTrasactionPembeli != nil {
	//	trxHandle.Rollback()
	//	res.Meta.Code = constants.BAD_REQUEST_CODE
	//	res.Meta.Status = false
	//	res.Meta.Message = errTransaction.Error()
	//	return
	//}
	//
	//// Save Balance transaction penjual
	//_, errBalanceTrasactionPenjual := svc.BalanceTransactionRepo.WithTrx(trxHandle).Save(balanceTransactionPenjual)
	//if errBalanceTrasactionPenjual != nil {
	//	trxHandle.Rollback()
	//	res.Meta.Code = constants.BAD_REQUEST_CODE
	//	res.Meta.Status = false
	//	res.Meta.Message = errTransaction.Error()
	//	return
	//}

	// Set Data for Customer Property Lot New
	//customerPropertyLotNew := models.CustomerPropertyLots{
	//	CustomerID:          int(customer.ID),
	//	PropertyID:          int(property.ID),
	//	Lot:                 req.Lot,
	//	TransactionDetailID: resTransaction.TransactionDetail.ID,
	//}
	//
	//// Save Customer Property Lot Pembeli
	//_, errCustomerPropertyLot := svc.CustomerPropertyLotRepo.WithTrx(trxHandle).Save(customerPropertyLotNew)
	//if errCustomerPropertyLot != nil {
	//	trxHandle.Rollback()
	//	res.Meta.Code = constants.BAD_REQUEST_CODE
	//	res.Meta.Status = false
	//	res.Meta.Message = errTransaction.Error()
	//	return
	//}

	// Update Customer Property Lot Penjual
	//customerPropertyLot.Lot = customerPropertyLot.Lot - req.Lot
	//_, errCustomerPropertyLot = svc.CustomerPropertyLotRepo.WithTrx(trxHandle).Save(customerPropertyLot)
	//if errCustomerPropertyLot != nil {
	//	trxHandle.Rollback()
	//	res.Meta.Code = constants.BAD_REQUEST_CODE
	//	res.Meta.Status = false
	//	res.Meta.Message = errTransaction.Error()
	//	return
	//}

	if paymentMethod.Code == "saldo" {
		//Update Customer Property Secondary Penjual
		customerPropertySecondary.Status = constants.SECONDARY_STATUS_CLOSE
		_, errCustomerPropertySecondary := svc.CustomerPropertySecondaryRepo.WithTrx(trxHandle).Save(customerPropertySecondary)
		if errCustomerPropertySecondary != nil {
			trxHandle.Rollback()
			res.Meta.Code = constants.BAD_REQUEST_CODE
			res.Meta.Status = false
			res.Meta.Message = errTransaction.Error()
			return
		}

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
			// Set Data for Customer Property Lot New
			customerPropertyLotNew := models.CustomerPropertyLots{
				CustomerID:          int(customer.ID),
				PropertyID:          int(property.ID),
				Lot:                 req.Lot,
				TransactionDetailID: resTransaction.TransactionDetail.ID,
			}
			//
			//// Save Customer Property Lot Pembeli
			errCustomerPropertyLot := svc.CustomerPropertyLotRepo.WithTrx(trxHandle).UpdateCustomerPropertyLot(customerPropertyLotNew)
			if errCustomerPropertyLot != nil {
				trxHandle.Rollback()
				res.Meta.Code = constants.BAD_REQUEST_CODE
				res.Meta.Status = false
				res.Meta.Message = errTransaction.Error()
				return
			}
		} else {
			existingCustomerPropertyLot.Lot = req.Lot + existingCustomerPropertyLot.Lot
			//// Save Customer Property Lot Pembeli
			errCustomerPropertyLot := svc.CustomerPropertyLotRepo.WithTrx(trxHandle).UpdateCustomerPropertyLot(existingCustomerPropertyLot)
			if errCustomerPropertyLot != nil {
				trxHandle.Rollback()
				res.Meta.Code = constants.BAD_REQUEST_CODE
				res.Meta.Status = false
				res.Meta.Message = errTransaction.Error()
				return
			}
		}

		existingCustomerPropertyLotParent, err := svc.CustomerPropertyLotRepo.FindByCustomerIDAndPropertyId(int(customerPropertyLot.CustomerID), req.PropertyID)
		if err != nil {
			trxHandle.Rollback()
			res.Meta.Code = constants.BAD_REQUEST_CODE
			res.Meta.Status = false
			res.Meta.Message = "Silahkan hubungi administrator"
			return
		}
		// Update Customer Property Lot Penjual
		existingCustomerPropertyLotParent.Lot = existingCustomerPropertyLotParent.Lot - req.Lot
		errCustomerPropertyLotSeller := svc.CustomerPropertyLotRepo.WithTrx(trxHandle).UpdateCustomerPropertyLot(customerPropertyLot)
		if errCustomerPropertyLotSeller != nil {
			trxHandle.Rollback()
			res.Meta.Code = constants.BAD_REQUEST_CODE
			res.Meta.Status = false
			res.Meta.Message = errTransaction.Error()
			return
		}

		fee := resTransaction.SubTotal * 0.005
		parentSubBalance := resTransaction.SubTotal - fee

		childCustomer, _ := svc.CustomerRepo.FindByID(uint64(customer.ID))
		if int(customer.ID) == customerPropertySecondary.CustomerID {
			fmt.Println("child", childCustomer.BalanceAmount)
			fmt.Println("parentSubBalance", parentSubBalance)
			fmt.Println("grandTotal", grandTotal)
			childCustomer.BalanceAmount = (childCustomer.BalanceAmount - grandTotal) + parentSubBalance
		} else {
			childCustomer.BalanceAmount = childCustomer.BalanceAmount - grandTotal
		}
		_, errUpdateCustomer := svc.CustomerRepo.WithTrx(trxHandle).UpdateBalance(childCustomer)
		if errUpdateCustomer != nil {
			trxHandle.Rollback()
			res.Meta.Code = constants.BAD_REQUEST_CODE
			res.Meta.Status = false
			res.Meta.Message = "Error, silahkan hubungi administrator"
			return
		}

		parentCustomer, _ := svc.CustomerRepo.FindByID(uint64(customerPropertySecondary.CustomerID))

		//Set Data for Balance Transaction
		balanceTransactionParent := models.BalanceTransaction{
			CustomerID:      int(parentCustomer.ID),
			TransactionType: constants.PENJUALAN_LOT,
			Description:     property.Name,
			Amount:          parentSubBalance,
			Status:          constants.STATUS_DONE,
			Date:            time.Now(),
			Note:            "",
			Lot:             req.Lot,
		}
		// Save Balance transaction
		_, errBalanceTrasactionParent := svc.BalanceTransactionRepo.WithTrx(trxHandle).Save(balanceTransactionParent)
		if errBalanceTrasactionParent != nil {
			trxHandle.Rollback()
			res.Meta.Code = constants.BAD_REQUEST_CODE
			res.Meta.Status = false
			res.Meta.Message = "Error, silahkan hubungi administrator"
			return
		}
		if int(customer.ID) != customerPropertySecondary.CustomerID {
			parentCustomer.BalanceAmount = parentCustomer.BalanceAmount + parentSubBalance
			_, errUpdateCustomerParent := svc.CustomerRepo.WithTrx(trxHandle).UpdateBalance(parentCustomer)
			if errUpdateCustomerParent != nil {
				trxHandle.Rollback()
				res.Meta.Code = constants.BAD_REQUEST_CODE
				res.Meta.Status = false
				res.Meta.Message = "Error, silahkan hubungi administrator"
				return
			}
		}

		//Set Data for Transaction Detail
		parentTransactionDetail := models.TransactionDetails{
			PropertyID: int(property.ID),
			Lot:        req.Lot,
			Price:      customerPropertySecondary.PricePerLot,
		}

		parentTransaction := models.Transactions{
			CustomerID:        customerPropertySecondary.CustomerID,
			SubTotal:          resTransaction.SubTotal,
			GrandTotal:        resTransaction.SubTotal,
			Fee:               fee,
			TransactionDetail: parentTransactionDetail,
			StatusId:          status,
			PaymentMethodId:   paymentMethod.ID,
			NoTransaction:     noTransction,
			TransactionType:   constants.PENJUALAN_LOT,
		}

		_, errTransactionParent := svc.TransactionRepo.WithTrx(trxHandle).Save(parentTransaction)
		if errTransactionParent != nil {
			trxHandle.Rollback()
			res.Meta.Code = constants.BAD_REQUEST_CODE
			res.Meta.Status = false
			res.Meta.Message = "Silahkan hubungi administrator"
			return
		}
	} else {
		//Update Customer Property Secondary Penjual
		customerPropertySecondary.Status = constants.SECONDARY_STATUS_PENDING
		_, errCustomerPropertySecondary := svc.CustomerPropertySecondaryRepo.WithTrx(trxHandle).Save(customerPropertySecondary)
		if errCustomerPropertySecondary != nil {
			trxHandle.Rollback()
			res.Meta.Code = constants.BAD_REQUEST_CODE
			res.Meta.Status = false
			res.Meta.Message = errTransaction.Error()
			return
		}
	}

	// End of transaction
	trxHandle.Commit()

	//Check to Mutation Bank BCA
	checkMutationBank()

	res.Data = token
	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
}
