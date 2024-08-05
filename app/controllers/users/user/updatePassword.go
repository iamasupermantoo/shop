package user

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

const updatePassword = 1
const updateSecurityPassword = 2

type UpdatePasswordParams struct {
	Type        int    `json:"type" validate:"required,oneof=1 2"`
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

// UpdatePassword 更新用户密码｜密钥
func UpdatePassword(ctx *context.CustomCtx, params *UpdatePasswordParams) error {
	userInfo := usersModel.User{}
	result := database.Db.Model(&userInfo).Where("id = ?", ctx.UserId).Find(&userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation", "updatePassword")
	}

	updateField := ""
	switch params.Type {
	case updatePassword:
		if userInfo.Password != utils.PasswordEncrypt(params.OldPassword) {
			return ctx.ErrorJsonTranslate("theOldPasswordIsIncorrect")
		}
		updateField = "password"
	case updateSecurityPassword:
		if userInfo.SecurityKey != utils.PasswordEncrypt(params.OldPassword) {
			return ctx.ErrorJsonTranslate("theOldPasswordIsIncorrect")
		}
		updateField = "security_key"
	}

	database.Db.Model(&usersModel.User{}).Where("id = ?", userInfo.ID).Update(updateField, utils.PasswordEncrypt(params.NewPassword))
	return ctx.SuccessJsonOK()
}
