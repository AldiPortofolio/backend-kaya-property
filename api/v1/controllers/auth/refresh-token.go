package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"
	"kaya-backend/models/request"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c authontroller) RefreshToken(ctx *gin.Context) {
	fmt.Println(">>> RefreshToken - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	req := request.RefreshToken{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("API - RefreshToken (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.AuthService.RefreshToken(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - RefreshToken (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
