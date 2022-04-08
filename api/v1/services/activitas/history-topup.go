package activitas

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
	"math"
)

func (svc *ActivitasService) HistoryTopup(req request.ActivitasFilter, res *models.ResponsePagination) {
	data, total, err := svc.balanceTransactionRepo.HistoryTopup(req)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	pagination := models.MetaPagination{
		CurrentPage: int64(req.Page),
		NextPage:    int64(req.Page + 1),
		PrevPage:    int64(req.Page - 1),
		TotalPages:  int64(math.Ceil(float64(int64(total)) / float64(25))),
		TotalCount:  int64(total),
	}

	res.Data = data
	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "success"
	res.Pagination = pagination
}
