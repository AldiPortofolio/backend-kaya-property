package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"
	"kaya-backend/utils/helper"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c authontroller) Me(ctx *gin.Context) {
	fmt.Println(">>> SendVerification - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	tokenAuth, err := helper.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	customerID, err := helper.FetchAuth(tokenAuth)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	reqBytes, _ := json.Marshal(customerID)
	log.Info("API - SendVerification (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.CustomerService.Me(uint(customerID), &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - SendVerification (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
