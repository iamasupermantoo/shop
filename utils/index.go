package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/goccy/go-json"
)

// IsNumeric 判断是否数字
func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case string:
		return false
	}
	return true
}

// PasswordEncrypt 密码加密
func PasswordEncrypt(password string) string {
	if password == "" {
		return ""
	}

	has1 := md5.Sum([]byte(password))
	has2 := md5.Sum([]byte(fmt.Sprintf("%x", has1)))
	return fmt.Sprintf("%x", has2)
}

// StructSign 对象签名
func StructSign(params interface{}, key string) string {
	dataBytes, _ := json.Marshal(params)
	dataMap := map[string]interface{}{}
	_ = json.Unmarshal(dataBytes, &dataMap)
	dataMap["sign"] = key
	dataMapBytes, _ := json.Marshal(dataMap)
	return PasswordEncrypt(string(dataMapBytes))
}
