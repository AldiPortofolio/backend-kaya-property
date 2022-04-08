package customerpropertysecondary

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
)

func (svc *CustomerPropertySecondaryService) SumLotPropertySecondary(req request.FilterPropertySecondary, res *models.Response) {
	var (
		totalLot int
		nilaiSecondary float64
		nilaiProperty float64
		profit float64
	)
	datas, err := svc.CustomerPropertySecondaryRepo.ListByCustomer(req)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	for _, data := range datas {
			totalLot += data.Lot
			nilaiSecondary += data.PricePerLot
			nilaiProperty += data.Property.PricePerLot
	}
	profit = nilaiSecondary - nilaiProperty

	res.Data = map[string]interface{}{
		"total_lot":    totalLot,
		"nilai":        nilaiSecondary,
		"total_profit": profit,
	}
	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
}
