package dashboard

import (
	"kaya-backend/models"
	"kaya-backend/utils/constants"
)

func (svc *DashboardService) Detail(customerID int, res *models.Response) {
	var (
		nilaiAsset     float64
		totalLot       int
		totalLotSold   int
		totalValue     float64
		totalInvestasi float64
		profit         float64
		// totalPriceLot  float64
	)

	//Get Data Customer
	customer, err := svc.CustomerRepo.FindByID(uint64(customerID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	//Get Customer Property Lot
	assetLot, err := svc.CustomerPropertyLotRepo.FindByCustomerID(int(customer.ID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	for _, obj := range assetLot {

		if obj.Property.IsSold {
			totalLotSold += obj.Lot
			totalValue += obj.Property.PricePerLot * float64(obj.Lot)

			grossProfit := obj.Property.SoldPrice - obj.Property.Price
			feeTotal := 0;
			for _, fee := range obj.Property.PropertyFee{
				feeTotal += int(fee.Amount)
			}
			netProfit := grossProfit - float64(feeTotal)
			netProfitPerLot := netProfit / float64(obj.Property.Lot)
			profit += netProfitPerLot * float64(obj.Lot)
		} else{
			nilaiAsset += obj.Property.PricePerLot * float64(obj.Lot)
			totalLot += obj.Lot
			totalInvestasi += obj.Property.PricePerLot * float64(obj.Lot)
		}
	}

	// //Get Customer Property Secondary
	// req := request.FilterProperty{
	// 	CustomerID: int(customer.ID),
	// 	Status:     "CLOSED",
	// }
	// assetSecondary, _, err := svc.CustomerPropertySecondaryRepo.GetAll(req)
	// if err != nil {
	// 	res.Meta.Code = constants.BAD_REQUEST_CODE
	// 	res.Meta.Status = false
	// 	res.Meta.Message = err.Error()
	// 	return
	// }

	// for _, obj := range assetSecondary {
	// 	totalLotSold += obj.Lot
	// 	totalValue += obj.Property.PricePerLot * float64(obj.Lot)
	// 	totalPriceLot += obj.PricePerLot * float64(obj.Lot)
	// }
	// profit += totalPriceLot - totalValue

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
	res.Data = map[string]interface{}{
		"nilai_asset":     nilaiAsset,
		"saldo":           customer.BalanceAmount,
		"total_lot_aktif": totalLot,
		"total_investasi": totalInvestasi,
		"total_lot_sold":  totalLotSold,
		"total_value":     totalValue,
		"profit":          profit,
	}
}
