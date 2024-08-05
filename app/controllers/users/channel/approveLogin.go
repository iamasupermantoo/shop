package channel

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/usersService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"time"
)

type ApproveLoginParams struct {
	ID uint `validate:"required" json:"id"`
}

// ApproveLogin 授权登录
func ApproveLogin(ctx *context.CustomCtx, params *ApproveLoginParams) error {
	channelInfo := &usersModel.Channel{}
	result := database.Db.Model(channelInfo).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Where("mode = ?", usersModel.ChannelModeChannel).Where("status = ?", usersModel.ChannelStatusActive).
		Where("id = ?", params.ID).Find(channelInfo)
	if result.Error != nil || channelInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation", "NotFound ChannelInfo")
	}

	// 获取当前用户信息
	userInfo := usersModel.User{}
	result = database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation", "NotFound userInfo")
	}

	dataParams := &usersService.ApproveLoginParams{Symbol: channelInfo.Symbol, User: userInfo.UserName, Pass: userInfo.Password, Time: time.Now().Unix()}
	routeLink, err := usersService.NewUserChannel(channelInfo).ApproveLogin(dataParams)
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}

	return ctx.SuccessJson(routeLink)
}
