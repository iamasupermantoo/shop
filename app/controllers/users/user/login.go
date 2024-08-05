package user

import (
	"github.com/dchest/captcha"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/commonService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

type LoginParams struct {
	UserName   string `json:"userName" validate:"required"` //	用户名
	Password   string `json:"password" validate:"required"` //	密码
	CaptchaId  string `json:"captchaId"`
	CaptchaVal string `json:"captchaVal"`
}

// Login 用户登录
func Login(ctx *context.CustomCtx, params *LoginParams) error {
	adminSettingService := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId)
	loginTemplate := adminSettingService.CheckBoxToMaps("loginTemplate")
	freezeTemplate := adminSettingService.CheckBoxToMaps("freezeTemplate")

	// 如果需要验证码
	if loginTemplate["showVerify"] {
		if !captcha.VerifyString(params.CaptchaId, params.CaptchaVal) {
			return ctx.ErrorJsonTranslate("codeError")
		}
	}

	// 获取用户信息
	userInfo := &homeUserInfo{}
	database.Db.Model(userInfo).Where("user.user_name = ?", params.UserName).
		Where("user.status = ?", usersModel.UserStatusActive).
		Preload("AuthInfo").Preload("LevelInfo").Find(userInfo)
	if userInfo.ID == 0 || userInfo.Password != utils.PasswordEncrypt(params.Password) ||
		utils.ArrayUintIndexOf(ctx.GetAdminChildIds(), userInfo.AdminId) == -1 ||
		(userInfo.Status == usersModel.UserStatusDisable && freezeTemplate["closeLogin"]) {
		return ctx.ErrorJsonTranslate("usernameOrPasswordIncorrect")
	}

	return ctx.SuccessJson(&authData{
		UserInfo: userInfo,
		Token:    commonService.NewServiceToken(ctx.Rds).GenerateHomeToken(userInfo.AdminId, userInfo.ID),
	})
}

type authData struct {
	Token    string        `json:"token"`
	UserInfo *homeUserInfo `json:"userInfo"`
}
