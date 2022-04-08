package customer

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"kaya-backend/library/logger/v2"
	"kaya-backend/models"
	"net/http"
)

func (c customerController) SaveGuest(ctx *gin.Context) {
	fmt.Println(">>> SaveGuest - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	req := models.Guest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	fmt.Println("req", req)

	reqBytes, _ := json.Marshal(req)
	log.Info("API - Customer (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.CustomerService.SaveGuest(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Customer (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
