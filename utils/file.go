package utils

import (
	"errors"
	"os"
)

// PathExists 判断路径是否存在
func PathExists(path string) bool {
	//os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// PathMkdirAll 递归创建文件夹
func PathMkdirAll(filePath string) {
	if !PathExists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

// FileWrite 文件如果不存在那么创建,存在那么写入
func FileWrite(filePath string, writeContent []byte) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(filePath)
		if err != nil {
			return err
		}
	}

	_, err = file.Write(writeContent)
	if err != nil {
		return err
	}
	return nil
}
