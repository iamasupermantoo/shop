package invite

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

// CreateParams 新增参数
type CreateParams struct {
	AdminName string `json:"adminName"`
	UserName  string `json:"userName"`
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	var inviteAdminId uint
	var inviteUserId uint
	if params.AdminName != "" {
		adminInfo := &adminsModel.AdminUser{}
		result := database.Db.Model(adminInfo).Where("user_name = ?", params.AdminName).Where("parent_id IN ?", ctx.GetAdminChildIds()).Find(adminInfo)
		if result.Error != nil || adminInfo.ID == 0 {
			return ctx.ErrorJson("没有当前管理信息")
		}
		inviteAdminId = adminInfo.ID
	} else {
		userInfo := &usersModel.User{}
		result := database.Db.Model(userInfo).Where("user_name = ?", params.UserName).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(userInfo)
		if result.Error != nil || userInfo.ID == 0 {
			return ctx.ErrorJson("没有当前用户信息")
		}
		inviteAdminId = userInfo.AdminId
		inviteUserId = userInfo.ID
	}

	inviteInfo := &usersModel.Invite{}
	result := database.Db.Model(inviteInfo).Where("admin_id = ?", inviteAdminId).Where("user_id = ?", inviteUserId).Find(inviteInfo)
	if inviteInfo.ID > 0 {
		return ctx.ErrorJson("当前管理｜用户 验证码已经存在")
	}

	createInfo := &usersModel.Invite{
		AdminId: inviteAdminId,
		UserId:  inviteUserId,
		Code:    utils.NewRandom().SetNumberRunes().String(6),
	}
	result = database.Db.Create(createInfo)
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
