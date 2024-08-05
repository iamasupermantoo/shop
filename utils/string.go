package utils

import (
	"bytes"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// ToHump 转换驼峰方法
func ToHump(s string) string {
	sl := strings.Split(s, "_")

	ns := ""
	for i := 0; i < len(sl); i++ {
		ns += strings.ToUpper(sl[i][:1]) + sl[i][1:]
	}
	return ns
}

// ToUnderlinedWords 转换下划线单词
func ToUnderlinedWords(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}

			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}

// GenerateRandomString 随机生成指定长度字符串
func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// CamelToSnake 将驼峰命名转换为下划线命名
func CamelToSnake(s string) string {
	var buf bytes.Buffer
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

// StringToIntList 字符串转int数组
// 例子1：2，34，53 -> [2,34,53]
// 例子1：2-34-53 -> [2,34,53]
// 例子1：2&34&53 -> [2,34,53]
func StringToIntList(s string) []int {
	var intList []int
	var tmp []rune

	for _, v := range s {
		if !unicode.IsDigit(v) {
			if len(tmp) > 0 {
				parseInt, err := strconv.Atoi(string(tmp))
				if err == nil {
					intList = append(intList, parseInt)
				}
				tmp = tmp[:0] // 清空tmp
			}
		} else {
			tmp = append(tmp, v)
		}
	}

	if len(tmp) > 0 {
		parseInt, err := strconv.Atoi(string(tmp))
		if err == nil {
			intList = append(intList, parseInt)
		}
	}

	return intList
}
