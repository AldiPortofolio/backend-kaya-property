package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"
	"kaya-backend/models/request"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c authontroller) SendEmailResetPassword(ctx *gin.Context) {
	fmt.Println(">>> SendEmailResetPassword - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	req := request.ByEmail{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("API - SendEmailResetPassword (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.AuthService.SendEmailResetPassword(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - SendEmailResetPassword (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}

func (c authontroller) ResetPassword(ctx *gin.Context) {
	fmt.Println(">>> ResetPassword - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	req := request.ResetPassword{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("API - ResetPassword (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.AuthService.ResetPassword(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - ResetPassword (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
