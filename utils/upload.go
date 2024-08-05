package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"time"
)

const (
	assetsPath     = "/public"
	UploadFilePath = "/uploads"      //	上传文件路径
	MaxUploadSize  = 5 * 1024 * 1000 //	最大上传大小
)

// Upload 上传文件
func Upload(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 检查文件大小
	if file.Size > MaxUploadSize {
		return "", errors.New(fmt.Sprintf("File size limit < %v", MaxUploadSize))
	}

	//生成上传文件路径
	dir, _ := os.Getwd()
	filenameWithSuffix := path.Base(file.Filename)
	saveFileDir := UploadFilePath + "/" + time.Now().Format("200601")
	saveFileName := "/" + GenerateRandomString(12) + path.Ext(filenameWithSuffix)

	//查询指定的目录是否存在
	if _, err = os.Stat(dir + assetsPath + saveFileDir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir+assetsPath+saveFileDir, os.ModePerm)
	}

	saveFilePath := dir + assetsPath + saveFileDir + saveFileName
	dst, err := os.Create(saveFilePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, _ = io.Copy(dst, src)
	return saveFileDir + saveFileName, nil
}
