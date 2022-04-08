package banner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c bannerontroller) Banner(ctx *gin.Context) {
	fmt.Println(">>> Banner - Controller <<<")
	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	c.BannerService.Banner(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - Login (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
