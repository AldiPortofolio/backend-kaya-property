package properties

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
)

func (svc *PropertyService) GetCountPortofolio(req request.FilterPortfolio, res *models.Response) {
	var (
		totalLot    int
		nilai       float64
		totalProfit float64
	)

	//Get Customer Property Lot
	assetLot, err := svc.customerPropertyLotRepo.FindByFilterProperty(req)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	for _, obj := range assetLot {
		totalLot += obj.Lot
		nilai += obj.Property.PricePerLot * float64(obj.Lot)
		
		if obj.Property.IsSold {
			grossProfit := obj.Property.SoldPrice - obj.Property.Price
			feeTotal := 0;
			for _, fee := range obj.Property.PropertyFee{
				feeTotal += int(fee.Amount)
			}
			netProfit := grossProfit - float64(feeTotal)
			netProfitPerLot := netProfit / float64(obj.Property.Lot)
			totalProfit += netProfitPerLot * float64(obj.Lot)
		}
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
	res.Data = map[string]interface{}{
		"total_lot":    totalLot,
		"nilai":        nilai,
		"total_profit": totalProfit,
	}
}
