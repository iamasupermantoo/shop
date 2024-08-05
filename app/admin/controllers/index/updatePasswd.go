package index

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

const (
	updatePasswdTypePasswd   = 1
	updatePasswdTypeSecurity = 2
)

// UpdatePasswdParams 更新密码参数
type UpdatePasswdParams struct {
	Type        int    `json:"type" validate:"required,oneof=1 2"` //	更新类型 1登录密码 2安全密码
	OldPassword string `json:"oldPassword" validate:"required"`    //	旧密码
	NewPassword string `json:"newPassword" validate:"required"`    //	新密码
	CmfPassword string `json:"cmfPassword" validate:"required"`    //	确认密码
}

// UpdatePasswd 更新密码
func UpdatePasswd(ctx *context.CustomCtx, params *UpdatePasswdParams) error {
	adminInfo := &adminsModel.AdminUser{}
	result := database.Db.Model(adminInfo).Where("id = ?", ctx.AdminId).Find(adminInfo)
	if result.Error != nil || adminInfo.ID == 0 {
		return ctx.ErrorJson("找不到管理信息")
	}

	fieldName := ""
	switch params.Type {
	case updatePasswdTypePasswd:
		if adminInfo.Password != utils.PasswordEncrypt(params.OldPassword) {
			return ctx.ErrorJson("旧登录密码不对")
		}
		fieldName = "password"
	case updatePasswdTypeSecurity:
		if adminInfo.SecurityKey != utils.PasswordEncrypt(params.OldPassword) {
			return ctx.ErrorJson("旧安全密码不对")
		}
		fieldName = "security_key"
	}

	if err := database.Db.Model(&adminsModel.AdminUser{}).Where("id = ?", adminInfo.ID).Update(fieldName, utils.PasswordEncrypt(params.NewPassword)).Error; err != nil {
		return ctx.ErrorJson(err.Error())
	}
	return ctx.SuccessJsonOK()
}
