package helper

import (
	"fmt"
	"kaya-backend/models/request"
	"kaya-backend/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	redisKaya "kaya-backend/utils/redis"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/twinj/uuid"
)

func CreateToken(customerID uint) (token request.Token, err error) {
	secretKey := utils.GetEnv("SECRET_KEY", "jwye99wbjskayafkjkdjserviceskdjkdrnmwrt")
	REFRESH_SECRET_KEY := utils.GetEnv("REFRESH_SECRET_KEY", "jwye99wbjskayafkjkdjserviceskdsmdnssnrt")

	td := request.Token{}
	td.AtExpires = time.Now().Add(time.Hour * 1).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["customer_id"] = customerID
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	td.AccessToken, err = at.SignedString([]byte(secretKey))
	if err != nil {
		return token, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["customer_id"] = customerID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(REFRESH_SECRET_KEY))
	if err != nil {
		return token, err
	}

	return td, nil
}

func CreateAuth(customerID int, td *request.Token) error {
	client := redisKaya.RedisStore()
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(td.AccessUuid, strconv.Itoa(int(customerID)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := client.Set(td.RefreshUuid, strconv.Itoa(int(customerID)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	secretKey := utils.GetEnv("SECRET_KEY", "jwye99wbjskayafkjkdjserviceskdjkdrnmwrt")
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*request.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		customerID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["customer_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &request.AccessDetails{
			AccessUuid: accessUuid,
			CustomerID: customerID,
		}, nil
	}
	return nil, err
}

func FetchAuth(authD *request.AccessDetails) (uint64, error) {
	client := redisKaya.RedisStore()
	customerid, err := client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	customerID, _ := strconv.ParseUint(customerid, 10, 64)
	return customerID, nil
}

func DeleteAuth(givenUuid string) (int64, error) {
	client := redisKaya.RedisStore()
	deleted, err := client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
