package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c dashboardController) DetailTicket(ctx *gin.Context) {
	fmt.Println(">>> Detail - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	reqID := ctx.Params.ByName("id")
	ID, _ := strconv.Atoi(reqID)

	res := models.Response{}

	reqBytes, _ := json.Marshal(ID)
	log.Info("API - Customer Verify (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.DashboardService.DetailTicket(ID, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Customer Verify (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
