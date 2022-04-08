package testimonial

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func (c testimonialController) GetAll(ctx *gin.Context) {
	fmt.Println(">>> GetAll - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	c.testimonialService.GetAll(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("API - GetAll (Response)", log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
