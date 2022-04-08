package auth

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"kaya-backend/utils"
	"kaya-backend/utils/constants"
	"kaya-backend/utils/helper"
	"strconv"

	"github.com/golang-jwt/jwt"
)

func (svc *ServiceAuth) RefreshToken(req request.RefreshToken, res *models.Response) {
	REFRESH_SECRET_KEY := utils.GetEnv("REFRESH_SECRET_KEY", "jwye99wbjskayafkjkdjserviceskdsmdnssnrt")
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(REFRESH_SECRET_KEY), nil
	})

	if err != nil {
		res.Meta.Code = constants.UNAUTHORIZED_CODE
		res.Meta.Status = false
		res.Meta.Message = "Token expired"
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		res.Meta.Code = constants.UNAUTHORIZED_CODE
		res.Meta.Status = false
		res.Meta.Message = "Token invalid"
		return
	}

	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			res.Meta.Code = constants.UNAUTHORIZED_CODE
			res.Meta.Status = false
			res.Meta.Message = "Refresh token invalid"
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["customer_id"]), 10, 64)
		if err != nil {
			res.Meta.Code = constants.UNAUTHORIZED_CODE
			res.Meta.Status = false
			res.Meta.Message = "Error occured"
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := helper.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			res.Meta.Code = constants.UNAUTHORIZED_CODE
			res.Meta.Status = false
			res.Meta.Message = "unauthorize"
			return
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := helper.CreateToken(uint(userId))
		if createErr != nil {
			res.Meta.Code = constants.UNAUTHORIZED_CODE
			res.Meta.Status = false
			res.Meta.Message = createErr.Error()
			return
		}
		//save the tokens metadata to redis
		saveErr := helper.CreateAuth(int(userId), &ts)
		if saveErr != nil {
			res.Meta.Code = constants.UNAUTHORIZED_CODE
			res.Meta.Status = false
			res.Meta.Message = saveErr.Error()
			return
		}
		tokens := map[string]interface{}{
			"access_token":         ts.AccessToken,
			"refresh_token":        ts.RefreshToken,
			"access_token_expires": ts.AtExpires,
		}

		res.Meta.Code = constants.SUCCESS_CODE
		res.Meta.Status = true
		res.Meta.Message = "Refresh token berhasil"
		res.Data = tokens
		return
	} else {
		res.Meta.Code = constants.UNAUTHORIZED_CODE
		res.Meta.Status = false
		res.Meta.Message = "Refresh token expired"
		return
	}
}
