package customer

import (
	"kaya-backend/models"
	"kaya-backend/utils/constants"
	"strings"
)

func (svc *CustomerService) SaveGuest(req models.Guest, res *models.Response) {
	_, err := svc.CustomerRepo.SaveGuest(req)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		if strings.Contains(err.Error(), "customer_uniq_email_and_no_handphone") {
			res.Meta.Message = "Email or No Hp already exist"
		} else {
			res.Meta.Message = "Error, Silahkan hubungi administrator"
		}

		return
	}

	res.Meta.Code = constants.CREATED_CODE
	res.Meta.Status = true
	res.Meta.Message = "Kami akan menghubungi anda melalui email dalam beberapa hari."

}
