package utils

import (
	"strings"

	"github.com/goccy/go-json"
)

// ArrayStringIndexOf 数组查找字符串是否存在
func ArrayStringIndexOf(arr []string, name string) int {
	for i, v := range arr {
		if v == name {
			return i
		}
	}
	return -1
}

// ArrayUintIndexOf 数组查询无符号是否存在
func ArrayUintIndexOf(arr []uint, id uint) int {
	for i, v := range arr {
		if v == id {
			return i
		}
	}
	return -1
}

// ArrayStringContainsOf 字符串模糊查找
func ArrayStringContainsOf(arr []string, name string) int {
	for i, v := range arr {
		if strings.Contains(v, name) {
			return i
		}
	}
	return -1
}

// StringToBoolMaps 值转换成 Bool maps
func StringToBoolMaps(value string) map[string]bool {
	dataList := make([]map[string]interface{}, 0)
	data := make(map[string]bool)
	_ = json.Unmarshal([]byte(value), &dataList)
	for _, v := range dataList {
		key, ok := v["field"].(string)
		val, ok1 := v["value"].(bool)
		if ok && ok1 {
			data[key] = val
		}
	}
	return data
}

// HttpGetParamsValues 参数转Get请求参数
func HttpGetParamsValues(values map[string]string) string {
	if values == nil {
		return ""
	}
	sl := make([]string, 0)
	for k, v := range values {
		if v != "" {
			sl = append(sl, k+"="+v)
		}
	}
	return "?" + strings.Join(sl, "&")
}
