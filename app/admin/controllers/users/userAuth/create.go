package userAuth

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	UserName string `json:"userName"` // 用户账户
	RealName string `json:"realName"` // 真实姓名
	Number   string `json:"number"`   // 卡号
	Photo1   string `json:"photo1"`   // 证件照1
	Photo2   string `json:"photo2"`   // 证件照2
	Type     int    `json:"type"`     // 类型 1身份证
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	userInfo := &usersModel.User{}
	result := database.Db.Model(userInfo).Where("user_name = ?", params.UserName).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJson("找不到当前用户")
	}
	createInfo := &usersModel.UserAuth{
		AdminId:  userInfo.AdminId,
		UserId:   userInfo.ID,
		RealName: params.RealName,
		Number:   params.Number,
		Photo1:   params.Photo1,
		Photo2:   params.Photo2,
		Type:     params.Type,
	}

	result = database.Db.Create(createInfo)
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
