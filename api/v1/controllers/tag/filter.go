package tag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"
	"kaya-backend/models/request"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c tagController)Filter(ctx *gin.Context) {
	fmt.Println(">>> tagFilter - Controller <<<")

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
	log.Info("API - Tag Filter (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.TagService.Filter(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Tag Filter (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusCreated, res)
}
