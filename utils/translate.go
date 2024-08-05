package utils

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// javascript "encodeURI()"
// so we embed js to our golang programm
func encodeURI(s string) string {
	return url.QueryEscape(s)
}

// Translate google 翻译文本
func Translate(source, sourceLang, targetLang string) (string, error) {
	var text []string
	var result []interface{}

	encodedSource := encodeURI(source)
	translateURL := "https://translate.google.com/translate_a/single?client=gtx&sl=" +
		sourceLang + "&tl=" + targetLang + "&dt=t&q=" + encodedSource

	r, err := http.Get(translateURL)
	if err != nil {
		return "err", errors.New("获取 translate.googleapis.com 时出错")
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "err", errors.New("读取响应正文时出错")
	}

	bReq := strings.Contains(string(body), `<title>Error 400 (Bad Request)`)
	if bReq {
		return "err", errors.New("错误 400（错误请求）")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "err", errors.New("解组数据时出错")
	}

	if len(result) > 0 {
		inner := result[0]
		for _, slice := range inner.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				text = append(text, fmt.Sprintf("%v", translatedText))
				break
			}
		}
		cText := strings.Join(text, "")

		return cText, nil
	} else {
		return "err", errors.New("没有回复翻译数据")
	}
}
