package dashboard

import (
	"kaya-backend/models"
	"kaya-backend/utils/constants"
)

func (svc *DashboardService) SaveTicket(req models.Ticket, res *models.Response) {
	req.Status = "OPEN"
	data, err := svc.ticketRepo.Save(req)
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
