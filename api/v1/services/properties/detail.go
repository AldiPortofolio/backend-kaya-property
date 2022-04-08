package properties

import (
	"kaya-backend/models"
	"kaya-backend/utils/constants"
)

func (svc *PropertyService) Detail(slug string, res *models.Response) {
	data, err := svc.propertyRepo.FindBySlug(slug)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	customerPropertyLot, err := svc.customerPropertyLotRepo.FindByPropertyID(int(data.ID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	totalLot := 0
	for _, obj := range customerPropertyLot {
		totalLot += obj.Lot
	}

	lotSold := data.PricePerLot * float64(totalLot)
	data.LotSold = lotSold

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
	res.Data = data
}
