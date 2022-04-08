package upload

import (
	"fmt"
	"net/http"

	"kaya-backend/models"
	"kaya-backend/utils/email"
	"kaya-backend/utils/helper"

	logger "kaya-backend/library/logger/v2"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	fmt.Println(">>> Upload - Controller <<<")

	log := logger.InitLogs(ctx.Request)

	res := models.Response{}

	fileName, err := helper.Upload(ctx)
	if err != nil {
		go log.Error(fmt.Sprintf("Upload error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Upload error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res.Meta.Code = 201
	res.Meta.Message = "Upload success"
	res.Meta.Status = true
	res.Data = map[string]interface{}{
		"file_name": fileName,
	}

	ctx.JSON(http.StatusCreated, res)
}

func Email(ctx *gin.Context) {

	to := "aldyarviansyah@gmail.com"
	subject := "Coba"
	bcc := []string{"aldyarviansyah@gmail.com"}
	data := struct {
		Name string
		URL  string
	}{
		Name: "Dhanush",
		URL:  "http://geektrust.in",
	}

	err := email.SendEmail(to, bcc, subject, data, "email-verifikasi.html")
	ctx.JSON(http.StatusCreated, err)
}
