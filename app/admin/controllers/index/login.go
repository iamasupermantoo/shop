package index

import (
	"github.com/dchest/captcha"
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/commonService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"time"
)

// LoginParams 登录参数
type LoginParams struct {
	UserName   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	CaptchaId  string `json:"captchaId" validate:"required"`
	CaptchaVal string `json:"captchaVal" validate:"required"`
}

func Login(ctx *context.CustomCtx, params *LoginParams) error {
	//	检查验证码是否正确
	if !captcha.VerifyString(params.CaptchaId, params.CaptchaVal) {
		return ctx.ErrorJson("验证码不正确")
	}

	//	检测管理是否存在， 并且密码是否匹配
	adminInfo := &adminsModel.AdminUser{}
	result := database.Db.Model(adminInfo).Where("user_name = ?", params.UserName).
		Where("status = ?", adminsModel.AdminUserStatusActive).
		Find(adminInfo)
	if result.Error != nil || adminInfo.ID == 0 || adminInfo.Password != utils.PasswordEncrypt(params.Password) {
		return ctx.ErrorJson("账户或密码不正确")
	}

	// 如果时间已经过期
	if adminInfo.ExpiredAt.Unix() < time.Now().Unix() {
		return ctx.ErrorJson("账户已过期~")
	}

	routerList := adminsService.NewAdminAuth(ctx.Rds, adminInfo.ID).GetRedisAdminRouterList()
	menuList := adminsService.NewAdminMenu(ctx.Rds, adminInfo.ID).GetRedisAdminMenuList()

	//	生成Token
	return ctx.SuccessJson(&resData{
		Token:      commonService.NewServiceToken(ctx.Rds).GenerateAdminToken(adminInfo.ID),
		Info:       adminInfo,
		RouterList: routerList,
		MenuList:   menuList,
	})
}

type resData struct {
	Token      string                       `json:"token"`      //	Token
	Info       *adminsModel.AdminUser       `json:"info"`       //	管理信息
	RouterList []string                     `json:"routerList"` //	路由列表
	MenuList   []*adminsModel.AdminMenuInfo `json:"menuList"`   //	管理菜单
}
