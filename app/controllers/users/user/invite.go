package user

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

// Invite 邀请信息
func Invite(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	inviteInfo := usersModel.Invite{}
	result := database.Db.Model(&inviteInfo).Where("user_id = ?", ctx.UserId).Find(&inviteInfo)
	if result.Error == nil && inviteInfo.ID == 0 {
		// 新增当前验证码
		inviteInfo = usersModel.Invite{AdminId: ctx.AdminId, UserId: ctx.UserId, Code: utils.NewRandom().SetNumberRunes().String(6)}
		database.Db.Create(&inviteInfo)
	}

	return ctx.SuccessJson(&inviteData{
		Code: inviteInfo.Code,
	})
}

type inviteData struct {
	Code string `json:"code"` // 邀请码
}
