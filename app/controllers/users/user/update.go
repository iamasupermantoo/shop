package user

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type UpdateParams struct {
	Avatar    string `json:"avatar"`    //	头像
	NickName  string `json:"nickName"`  // 昵称
	Email     string `json:"email"`     // 邮箱
	Telephone string `json:"telephone"` // 手机号码
	Sex       int    `json:"sex"`       // 性别
	Birthday  string `json:"birthday"`  // 生日
	Desc      string `json:"desc"`      // 个性签名
}

// Update 更新用户信息
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&usersModel.User{}).Where("id = ?", ctx.UserId).Updates(params)
	if result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}
	return ctx.SuccessJsonOK()
}
