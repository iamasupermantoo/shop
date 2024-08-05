package context

import (
	"fmt"
	"gofiber/app/models/service/systemsService"
)

// RespJson 返回信息
type RespJson struct {
	Code int         `json:"code"` //	错误代码
	Data interface{} `json:"data"` //	内容
	Msg  string      `json:"msg"`  //	错误信息
}

// ErrorJsonTranslate 输出翻译错误
func (_CustomCtx *CustomCtx) ErrorJsonTranslate(msg string, args ...any) error {
	translateService := systemsService.NewSystemTranslate(_CustomCtx.Rds, _CustomCtx.AdminId)
	msg = translateService.GetRedisAdminTranslateLangField(_CustomCtx.Lang, msg)
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	return _CustomCtx.JSON(&RespJson{Code: -1, Msg: msg})
}

// ErrorJsonTranslateMultiple 多字段翻译
func (_CustomCtx *CustomCtx) ErrorJsonTranslateMultiple(args ...string) error {
	translateService := systemsService.NewSystemTranslate(_CustomCtx.Rds, _CustomCtx.AdminId)
	msg := ""
	for _, arg := range args {
		msg += " " + translateService.GetRedisAdminTranslateLangField(_CustomCtx.Lang, arg)
	}
	return _CustomCtx.JSON(&RespJson{Code: -1, Msg: msg})
}

// ErrorJson 输出错误Json
func (_CustomCtx *CustomCtx) ErrorJson(msg string) error {
	return _CustomCtx.JSON(&RespJson{Code: -1, Msg: msg})
}

// SuccessJson 输出成功Json
func (_CustomCtx *CustomCtx) SuccessJson(data any) error {
	return _CustomCtx.JSON(&RespJson{Data: data})
}

// SuccessJsonOK 输出成功OK
func (_CustomCtx *CustomCtx) SuccessJsonOK() error {
	return _CustomCtx.JSON(&RespJson{Data: "ok"})
}

// IsErrorJson 是否是错误
func (_CustomCtx *CustomCtx) IsErrorJson(err error) error {
	if err != nil {
		return _CustomCtx.ErrorJson(err.Error())
	}
	return _CustomCtx.SuccessJsonOK()
}

// systemTranslate 系统错误翻译
func (_CustomCtx *CustomCtx) validateError(msg string, args ...any) error {
	translateService := systemsService.NewSystemTranslate(_CustomCtx.Rds, 0)
	msg = translateService.GetRedisAdminTranslateLangField(_CustomCtx.Lang, msg)
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	return _CustomCtx.JSON(&RespJson{Code: -1, Msg: msg})
}
