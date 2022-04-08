package dashboard

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"kaya-backend/library/logger/v2"
	"kaya-backend/models"
	"kaya-backend/utils/helper"
	"net/http"
	"strconv"
)

func (c dashboardController) CloseTicket(ctx *gin.Context) {
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
	reqID := ctx.Params.ByName("id")
	ID, _ := strconv.Atoi(reqID)

	reqBytes, _ := json.Marshal(customerID)
	log.Info("API - Customer Verify (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.DashboardService.CloseTicket(ID, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Customer Verify (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
