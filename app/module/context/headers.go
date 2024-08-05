package context

import (
	"github.com/golang-jwt/jwt/v5"
	"gofiber/app/config"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/commonService"
	"net/url"
	"strings"
)

func (_CustomCtx *CustomCtx) GetToken() string {
	userToken := _CustomCtx.Query("token")
	if userToken != "" {
		return userToken
	}

	// 头信息获取Token
	headers := _CustomCtx.GetReqHeaders()
	if _, ok := headers["Authorization"]; ok {
		if len(headers["Authorization"]) > 0 {
			authorizationList := strings.Split(headers["Authorization"][0], " ")
			if len(authorizationList) == 2 {
				return authorizationList[1]
			}
		}
	}
	return ""
}

// GetLang 获取当前请求语言
func (_CustomCtx *CustomCtx) initLang() {
	_CustomCtx.Lang = _CustomCtx.Query("lang")
	if _CustomCtx.Lang == "" {
		_CustomCtx.Lang = _CustomCtx.Get("Accept-Language")
		if _CustomCtx.Lang == "" {
			_CustomCtx.Lang = config.Conf.Lang
		}
	}
}

// GetOriginHost 获取来源Host
func (_CustomCtx *CustomCtx) initOriginHost() {
	originURL, err := url.Parse(_CustomCtx.Get("Origin"))
	if err == nil {
		_CustomCtx.OriginHost = originURL.Host
	}
}

// GetClaims 获取上下文 Claims
func (_CustomCtx *CustomCtx) initClaims() {
	adminService := adminsService.NewAdminUser(_CustomCtx.Rds, _CustomCtx.AdminId)
	if userToken, ok := _CustomCtx.Locals("token").(*jwt.Token); ok {
		claims := userToken.Claims.(jwt.MapClaims)
		if claims != nil {
			_CustomCtx.AdminId = uint(claims["adminId"].(float64))
			_CustomCtx.UserId = uint(claims["userId"].(float64))
			_CustomCtx.AdminSettingId = adminService.GetRedisAdminSettingId(_CustomCtx.AdminId)
			return
		}
	}

	//	获取Query Token 或者头信息 Token 来解析
	tokenStr := _CustomCtx.GetToken()
	_CustomCtx.AdminId, _CustomCtx.UserId = commonService.NewServiceToken(_CustomCtx.Rds).VerifyToken(_CustomCtx.Locals("ServiceName").(string), tokenStr)
	if _CustomCtx.AdminId > 0 {
		_CustomCtx.AdminSettingId = adminService.GetRedisAdminSettingId(_CustomCtx.AdminId)
	}

	// 如果都找不到, 那么使用源站域名获取管理ID
	if _CustomCtx.AdminId == 0 && _CustomCtx.UserId == 0 {
		_CustomCtx.AdminId = adminService.GetRedisDomainAdminId(_CustomCtx.OriginHost)
		if _CustomCtx.AdminId > 0 {
			_CustomCtx.AdminSettingId = adminService.GetRedisAdminSettingId(_CustomCtx.AdminId)
		}
	}
}
