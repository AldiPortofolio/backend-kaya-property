package helper

import (
	"fmt"
	"kaya-backend/utils"
	"kaya-backend/utils/constants"
	httpUtils "kaya-backend/utils/http"
	"net/http"
)

func SendVerification(noHp string, verifyCode string) error {
	zenzivaOn := utils.GetEnv("ZENZIVA_ON", "0")
	zenzivaEnpoint := utils.GetEnv("ZENZIVA_ENDPOINT", "https://console.zenziva.net/wareguler/api/sendWA/")
	zenzivaUserKey := utils.GetEnv("ZENZIVA_USER_KEY", "d3d73f1714d3")
	zenzivaPassKey := utils.GetEnv("ZENZIVA_PASS_KEY", "0581b880f3373945c851e954")

	if zenzivaOn == "0" {
		return nil
	}

	header := make(http.Header)
	header.Add("Content-Type", "application/json")

	requestBody := map[string]string{
		"userkey": zenzivaUserKey,
		"passkey": zenzivaPassKey,
		"to":      noHp,
		"message": fmt.Sprintf("Kode verifikasi Anda adalah %s", verifyCode),
	}

	responseBody, err := httpUtils.SendHttpRequest(constants.HttpMethodPost, zenzivaEnpoint, header, requestBody)
	if err != nil {
		return err
	}
	fmt.Println("responseBody", responseBody)

	return nil
}
