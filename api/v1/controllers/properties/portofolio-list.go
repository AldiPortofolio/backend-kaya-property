package properties

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"
	"kaya-backend/models/request"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c propertiesController) ListPortolio(ctx *gin.Context) {
	fmt.Println(">>> List - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.ResponsePagination{}

	req := request.FilterProperty{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("API - Property (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.PropertiesService.GetAllPortfolio(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Property (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
