package dashboard

import (
	"encoding/json"
	"fmt"
	"kaya-backend/utils/constants"
	"net/http"

	"kaya-backend/models"
	"kaya-backend/utils/helper"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c dashboardController) SellLot(ctx *gin.Context) {
	fmt.Println(">>> Detail - Controller <<<")

	log := logger.InitLogs(ctx.Request)

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

	res := models.Response{}
	req := models.CustomerPropertySecondaries{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	req.CustomerID = int(customerID)

	reqBytes, _ := json.Marshal(customerID)
	log.Info("API - Customer Verify (Request)", log.AddField("RequestBody", string(reqBytes)))

	req.Status = constants.SECONDARY_STATUS_OPEN

	c.DashboardService.SellLot(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Customer Verify (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
