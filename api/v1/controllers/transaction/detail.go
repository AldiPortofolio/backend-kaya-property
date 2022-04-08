package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"kaya-backend/library/logger/v2"
	"kaya-backend/models"
	"net/http"
)

func (c *transactionController) Detail(ctx *gin.Context) {
	fmt.Println(">>> Transaction Detail - Controller <<<")

	log := logger.InitLogs(ctx.Request)
	res := models.Response{}

	noOrder := ctx.Params.ByName("noOrder")
	reqBytes, _ := json.Marshal(noOrder)
	log.Info("API - Property (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.TransactionService.Detail(noOrder, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Property (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)

}
