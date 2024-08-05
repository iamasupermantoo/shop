package channel

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/commonService"
	"gofiber/app/models/service/usersService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"strings"
)

// Approve 授权登录处理
func Approve(ctx *context.CustomCtx, params *usersService.ApproveLoginParams) error {
	// 判断渠道是否存在
	channelInfo := &usersModel.Channel{}
	result := database.Db.Model(channelInfo).Where("mode = ?", usersModel.ChannelModeApprove).
		Where("status = ?", usersModel.ChannelStatusActive).
		Where("symbol = ?", params.Symbol).Find(channelInfo)
	if result.Error != nil || channelInfo.ID == 0 {
		return ctx.ErrorJson("NotFound ChannelInfo")
	}

	sign := utils.StructSign(params, channelInfo.Pass)
	if params.Sign != sign {
		return ctx.ErrorJson("Signature Failed")
	}

	// 判断当前用户是否存在
	userInfo := &usersModel.User{}
	result = database.Db.Model(userInfo).Where("user_name = ?", params.Symbol+"_"+params.User).Find(userInfo)
	if result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	// 如果不存在当前用户, 那么创建当前用户
	if userInfo.ID == 0 {
		userInfo = &usersModel.User{
			ChannelId: channelInfo.ID, AdminId: channelInfo.AdminId,
			UserName: params.Symbol + "_" + params.User, Password: params.Pass,
			SecurityKey: params.Pass, Type: usersModel.UserTypeChannel,
		}
	}

	// 获取当前渠道管理对应的域名
	adminInfo := &adminsModel.AdminUser{}
	result = database.Db.Model(adminInfo).Where("id = ?", channelInfo.AdminId).Find(adminInfo)
	if result.Error != nil || adminInfo.ID == 0 {
		return ctx.ErrorJson("NotFound AdminInfo")
	}

	domainList := strings.Split(adminInfo.Domains, ",")
	if len(domainList) <= 0 {
		return ctx.ErrorJson("NotFound Domains")
	}

	tokenStr := commonService.NewServiceToken(ctx.Rds).GenerateHomeToken(userInfo.AdminId, userInfo.ID)
	domainURL := ctx.Protocol() + "://" + domainList[0] + "/channel/login?token=" + tokenStr
	return ctx.SuccessJson(domainURL)
}
