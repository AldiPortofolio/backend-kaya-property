package properties

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils/helper"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c propertiesController) ListPortofolioMe(ctx *gin.Context) {
	fmt.Println(">>> List - Controller <<<")

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
	req.CustomerID = int(customerID)
	if req.Status == "ALL" {
		req.Status = ""
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("API - Property (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.PropertiesService.GetPortfolioMe(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Property (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
