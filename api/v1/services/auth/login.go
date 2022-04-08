package auth

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/helper"
)

func (svc *ServiceAuth) Login(req request.Login, res *models.Response) {
	dataOutput, result, err := svc.AuthRepo.Login(req)

	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	if !result {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "Invalid Username Or Password"
		return
	}

	token, err := helper.CreateToken(dataOutput.ID)
	if err != nil {
		res.Meta.Code = constants.INTERNAL_SERVER_ERROR_CODE
		res.Meta.Status = false
		res.Meta.Message = "Gagal saat membuat token"
		return
	}

	saveErr := helper.CreateAuth(int(dataOutput.ID), &token)
	if saveErr != nil {
		res.Meta.Code = constants.INTERNAL_SERVER_ERROR_CODE
		res.Meta.Status = false
		res.Meta.Message = "Gagal saat membuat token"
		return
	}

	data := map[string]interface{}{
		"id":      dataOutput.ID,
		"name":      dataOutput.Name,
		"email":     dataOutput.Email,
		"no_hp":     dataOutput.NoHp,
		"is_active": dataOutput.IsActive,
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "Login Berhasil"
	res.Data = map[string]interface{}{
		"customer":             data,
		"access_token":         token.AccessToken,
		"refresh_token":        token.RefreshToken,
		"access_token_expires": token.AtExpires,
	}
}
