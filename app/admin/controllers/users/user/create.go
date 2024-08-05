package user

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

// CreateParams 新增参数
type CreateParams struct {
	UserName string `validate:"required" json:"userName"` // 用户名
	Password string `validate:"required" json:"password"` // 密码
	Type     int    `validate:"required" json:"type"`     // 类型 -1虚拟用户 1默认用户 10会员用户
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	params.Password = utils.PasswordEncrypt(params.Password)
	createInfo := &usersModel.User{
		AdminId:  ctx.AdminId,
		UserName: params.UserName,
		NickName: params.UserName,
		Password: params.Password,
		Type:     params.Type,
	}

	result := database.Db.Create(createInfo)
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
