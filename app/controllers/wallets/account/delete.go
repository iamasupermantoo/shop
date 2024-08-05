package account

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

type DeleteParams struct {
	ID          int    `json:"id" validate:"required"` // 提现账户ID
	SecurityKey string `json:"securityKey"`            // 安全密钥
}

// Delete 提现账户删除
func Delete(ctx *context.CustomCtx, params *DeleteParams) error {
	walletsTemplate := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId).CheckBoxToMaps("walletsTemplate")

	// 是否需要安全密钥
	if walletsTemplate["showAccountSecurityKey"] {
		userInfo := &usersModel.User{}
		database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Find(userInfo)
		if userInfo.SecurityKey != utils.PasswordEncrypt(params.SecurityKey) {
			return ctx.ErrorJsonTranslate("incorrectSecurityKey")
		}
	}

	// 是否可以删除
	if !walletsTemplate["showAccountDelete"] {
		return ctx.SuccessJsonOK()
	}

	// 删除当前数据
	database.Db.Where("user_id = ?", ctx.UserId).Where("id = ?", params.ID).Delete(&walletsModel.WalletUserAccount{})
	return ctx.SuccessJsonOK()
}
