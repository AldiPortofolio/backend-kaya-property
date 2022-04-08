package properties

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c propertiesController) DetailListSecondary(ctx *gin.Context) {
	fmt.Println(">>> DetailListSecondary - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	slug := ctx.Params.ByName("slug")

	reqBytes, _ := json.Marshal(slug)
	log.Info("API - Property (Request)", log.AddField("RequestBody", string(reqBytes)))

	c.CustomerPropertySecondaryService.DetailListSecondary(slug, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Property (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
