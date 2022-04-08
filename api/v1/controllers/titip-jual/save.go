package titipjual

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c titipJualController) Save(ctx *gin.Context) {
	fmt.Println(">>> Save - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	req := models.TitipJual{}
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
	log.Info("API - Save (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.TitipJualService.Save(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Save (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
