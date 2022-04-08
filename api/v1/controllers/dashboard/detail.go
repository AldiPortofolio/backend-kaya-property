package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"
	"kaya-backend/utils/helper"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c dashboardController) Detail(ctx *gin.Context) {
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

	reqBytes, _ := json.Marshal(customerID)
	log.Info("API - Customer Verify (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.DashboardService.Detail(int(customerID), &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Customer Verify (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
