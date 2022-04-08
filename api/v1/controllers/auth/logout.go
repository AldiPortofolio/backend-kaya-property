package auth

import (
	"kaya-backend/utils/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c authontroller) Logout(ctx *gin.Context) {
	auth, err := helper.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	deleted, delErr := helper.DeleteAuth(auth.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	ctx.JSON(http.StatusOK, "Successfully logged out")
}
