package minio

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	// "github.com/extrame/xls"
	"kaya-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func initMinio() *minio.Client {
	endpoint := utils.GetEnv("MINIO_ENDPOINT", "192.168.8.117:9000")
	accessKeyID := utils.GetEnv("MINIO_ACCESSKEY", "AKIAIOSFODNN7KAYAACCESSKEY")
	secretAccessKey := utils.GetEnv("MINIO_SECRETKEY", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYKAYASECRETKEY")
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}

func FileUploadMinio(context *gin.Context, fileName string, fileBuffer io.Reader, fileSize int64, ContentType string) (string, error) {
	minioClient := initMinio()
	bucketName := utils.GetEnv("MINIO_BUCKET", "kaya-dev")
	hostMino := utils.GetEnv("MINIO_FILE_PATH", "http://192.168.8.117:9000")

	info, err := minioClient.PutObject(context, bucketName, fileName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: ContentType})
	if err != nil {
		return "", err
	}

	fullURL := fmt.Sprintf("%s/%s/%s", hostMino, info.Bucket, fileName)
	return fullURL, nil
}

func FileToBase64(filePath string, deleteFile bool) (string, string, error) {
	fileContent, _ := os.Open(filePath)
	fileName := path.Base(filePath)

	reader := bufio.NewReader(fileContent)
	content, _ := ioutil.ReadAll(reader)

	encoded := base64.StdEncoding.EncodeToString(content)
	if deleteFile {
		err := os.Remove(filePath)
		if err != nil {
			return encoded, fileName, err
		}
	}

	// fmt.Println("ENCODED: " + encoded)
	return encoded, fileName, nil
}
