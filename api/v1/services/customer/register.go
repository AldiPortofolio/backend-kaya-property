package customer

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/helper"
	"strings"
)

func (svc *CustomerService) Register(req models.Customers, res *models.Response) {
	fmt.Println(">>> Service Customer - Register <<<")

	req.IsActive = false
	hashedPassword, errPass := helper.HashPassword(req.Password)
	if errPass != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = errPass.Error()
		return
	}

	req.Password = hashedPassword
	if req.NoHp[0:1] != "0" {
		req.NoHp = "0" + req.NoHp
	}

	level, err := svc.MembershipLevelRepo.GetLevelDefault()
	if err != nil {
		res.Meta.Code = constants.INTERNAL_SERVER_ERROR_CODE
		res.Meta.Status = false
		res.Meta.Message = "Gagal mendapatkan membership level"
		return
	}

	req.MembershipLevelID = level.ID

	result, err := svc.CustomerRepo.Save(req)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		if strings.Contains(err.Error(), "customer_uniq_email_and_no_handphone") {
			res.Meta.Message = "Email or No Hp already exist"
		} else {
			res.Meta.Message = err.Error()
		}

		return
	}

	token, err := helper.CreateToken(result.ID)
	if err != nil {
		res.Meta.Code = constants.INTERNAL_SERVER_ERROR_CODE
		res.Meta.Status = false
		res.Meta.Message = "Gagal saat membuat token"
		return
	}

	saveErr := helper.CreateAuth(int(result.ID), &token)
	if saveErr != nil {
		res.Meta.Code = constants.INTERNAL_SERVER_ERROR_CODE
		res.Meta.Status = false
		res.Meta.Message = "Gagal saat membuat token"
		return
	}

	data := map[string]interface{}{
		"id":        result.ID,
		"name":      result.Name,
		"email":     result.Email,
		"no_hp":     result.NoHp,
		"is_active": result.IsActive,
	}

	res.Meta.Code = constants.CREATED_CODE
	res.Meta.Status = true
	res.Meta.Message = "Registrasi Anda berhasil dan sedang diproses oleh admin. Terima kasih."
	res.Data = map[string]interface{}{
		"customer":              data,
		"access_token":          token.AccessToken,
		"refresh_token":         token.RefreshToken,
		"access_token_expires":  token.AtExpires,
		"refresh_token_expires": token.RtExpires,
	}
}
