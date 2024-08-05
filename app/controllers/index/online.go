package index

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"net/url"
	"strconv"
)

// Online 跳转到内置在线客服
func Online(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	adminInfo := &adminsModel.AdminUser{}
	database.Db.Model(adminInfo).Where("id = ?", ctx.AdminId).Find(adminInfo)

	var onlineURL *url.URL
	var v url.Values
	var err error

	// 获取域名设置的客服链接
	if adminInfo.Online != "" {
		onlineURL, err = url.Parse(adminInfo.Online)
		if err == nil {
			v = onlineURL.Query()
			v.Add("l", ctx.Lang)
		}
	}

	// 用户登录之后的操作
	if ctx.UserId > 0 {
		userInfo := &usersModel.User{}
		adminInfo = &adminsModel.AdminUser{}
		database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Find(userInfo)
		database.Db.Model(adminInfo).Where("id = ?", ctx.AdminId).Find(adminInfo)
		onlineURL, err = url.Parse(adminInfo.Online)
		if err == nil && adminInfo.Online != "" {
			v = onlineURL.Query()
			v.Add("l", ctx.Lang)
			v.Add("n", userInfo.UserName)
			v.Add("p", utils.PasswordEncrypt(userInfo.UserName+strconv.Itoa(int(userInfo.ID))))
		}
	}

	// 组成正式客服链接
	if onlineURL != nil {
		adminInfo.Online = onlineURL.Scheme + "://" + onlineURL.Hostname() + onlineURL.Path
		if onlineURL.Port() != "" {
			adminInfo.Online += ":" + onlineURL.Port()
		}
		adminInfo.Online += "?" + v.Encode()
	}

	return ctx.SuccessJson(&onlineData{
		Icon: adminInfo.Avatar,
		Link: adminInfo.Online,
	})
}

type onlineData struct {
	Icon string `json:"icon"` // 图标
	Link string `json:"link"` // 链接
}
