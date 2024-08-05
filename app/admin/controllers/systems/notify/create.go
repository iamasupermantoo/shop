package notify

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	UserName string `gorm:"-" validate:"required" json:"userName"` // 用户名
	Mode     int    `validate:"required" json:"mode"`              // 模式 1后台 11前台
	Name     string `validate:"required" json:"name"`              // 标题
	Content  string `validate:"required" json:"content"`           // 内容
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	switch params.Mode {
	// 新增管理通知
	case systemsModel.NotifyModeAdminMessage:
		// 获取管理员信息
		adminUserInfo := &adminsModel.AdminUser{}
		result := database.Db.Where("user_name = ?", params.UserName).Find(adminUserInfo)
		if result.Error != nil || adminUserInfo.ID == 0 {
			return ctx.ErrorJson("找不到管理信息")
		}

		result = database.Db.Create(&systemsModel.Notify{
			AdminId: adminUserInfo.ID,
			UserId:  0,
			Name:    params.Name,
			Mode:    params.Mode,
			Content: params.Content,
		})
		if result.Error != nil {
			return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
		}

	// 新增用户通知
	case systemsModel.NotifyModeHomeMessage:
		// 获取用户信息
		userInfo := &usersModel.User{}
		result := database.Db.Where("user_name = ?", params.UserName).Find(userInfo)
		if result.Error != nil || userInfo.ID == 0 {
			return ctx.ErrorJson("找不到用户信息")
		}

		result = database.Db.Create(&systemsModel.Notify{
			AdminId: userInfo.AdminId,
			UserId:  userInfo.ID,
			Name:    params.Name,
			Mode:    params.Mode,
			Content: params.Content,
		})
		if result.Error != nil {
			return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
		}
	}
	return ctx.SuccessJsonOK()
}
