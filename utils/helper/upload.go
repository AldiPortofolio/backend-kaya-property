package helper

import (
	"fmt"
	"kaya-backend/utils"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) (string, error) {
	fmt.Println(">>> Helper - Upload <<<")
	uploadDirectory := utils.GetEnv("UPLOAD_DIRECTORY", "fileupload")

	tokenAuth, err := ExtractTokenMetadata(ctx.Request)
	if err != nil {
		fmt.Println("tokenAuth", err)
		return "", err
	}

	customerID, err := FetchAuth(tokenAuth)
	if err != nil {
		return "", err
	}
	fmt.Println("customerID", customerID)

	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("File not found")
		return "", err
	}
	fileName := file.Filename

	newDirectory := uploadDirectory + "/" + strconv.Itoa(int(customerID))
	_, errs := os.Stat(newDirectory)
	if os.IsNotExist(errs) {
		errDir := os.MkdirAll(newDirectory, 0755)
		if errDir != nil {
			fmt.Println("Gagal membuat folder", errDir)
			return "", errs
		}
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Gagal membaca directory", dir)
		return "", err
	}

	fileLocation := filepath.Join(dir, newDirectory, fileName)
	if err := ctx.SaveUploadedFile(file, fileLocation); err != nil {
		fmt.Println("Gagal upload", dir)
		return "", err
	}

	return fileName, nil
}

//// Get Buffer from file
//fileBuffer, err := file.Open()
//
//if err != nil {
//	go log.Error(fmt.Sprintf("File error: %v", err))
//	res.Meta.Code = 400
//	res.Meta.Message = fmt.Sprintf("File error: %v", err)
//	res.Meta.Status = false
//	ctx.JSON(http.StatusBadRequest, res)
//	return
//}
//defer fileBuffer.Close()
//
//minioUpload, err := minio.FileUploadMinio(ctx, fileName, fileBuffer, fileSize, contentType)
//
//if err != nil {
//	go log.Error(fmt.Sprintf("Minio upload error: %v", err))
//	res.Meta.Code = 400
//	res.Meta.Message = fmt.Sprintf("Minio upload error: %v", err)
//	res.Meta.Status = false
//	ctx.JSON(http.StatusBadRequest, res)
//	return
//}
