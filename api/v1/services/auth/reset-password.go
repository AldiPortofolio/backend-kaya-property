package auth

import (
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/email"
	"kaya-backend/utils/helper"
	"strconv"
)

func (svc ServiceAuth) SendEmailResetPassword(req request.ByEmail, res *models.Response) {
	endpoint := utils.GetEnv("ENDPOINT_RESET_PASSWORD", "http://localhost:3000/reset-password/reset")
	customer, err := svc.AuthRepo.ByEmail(req)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = "Email tidak terdaftar disistem kami, silahkan untuk mendaftar terlebih dahulu"
		return
	}

	token := helper.GenerateRandomCode2(15)

	reqToken := models.Tokens{
		Token:      token,
		CustomerID: int(customer.ID),
		IsActive:   true,
	}

	tokenData, err := svc.TokenRepo.Save(reqToken)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	urlResetPass := endpoint + "?t=" + tokenData.Token + "&u=" + strconv.Itoa(tokenData.CustomerID)

	to := customer.Email
	subject := "Reset Password Kaya"
	bcc := []string{customer.Email}
	data := struct {
		Name string
		URL  string
	}{
		Name: customer.Name,
		URL:  urlResetPass,
	}

	errSendmail := email.SendEmail(to, bcc, subject, data, "email-verifikasi.html")
	if errSendmail != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "link reset password berhasil dikirim by email"

}

func (svc ServiceAuth) ResetPassword(req request.ResetPassword, res *models.Response) {
	token, err := svc.TokenRepo.FindByFilter(req)
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	customer, err := svc.CustomerRepo.FindByID(uint64(req.CustomerID))
	if err != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	hashedPassword, errPass := helper.HashPassword(req.Password)
	if errPass != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = errPass.Error()
		return
	}

	customer.Password = hashedPassword

	_, errSave := svc.CustomerRepo.Save(customer)
	if errSave != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = errSave.Error()
		return
	}

	token.IsActive = false
	_, errSaveToken := svc.TokenRepo.Save(token)
	if errSaveToken != nil {
		res.Meta.Code = constants.BAD_REQUEST_CODE
		res.Meta.Status = false
		res.Meta.Message = errSaveToken.Error()
		return
	}

	res.Meta.Code = constants.SUCCESS_CODE
	res.Meta.Status = true
	res.Meta.Message = "reset password berhasil, silahkan login kembali"

}
