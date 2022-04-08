package blog

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"
	"kaya-backend/models/request"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c blogController)Filter(ctx *gin.Context) {
	fmt.Println(">>> blogFilter - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	req := request.TagBlog{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("API - Blog Filter (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.BlogService.Filter(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Blog Filter (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusCreated, res)
}
