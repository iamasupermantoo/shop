package account

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

type UpdateParams struct {
	ID          int    `json:"id" validate:"required" gorm:"-"` // 提现账户ID
	SecurityKey string ` json:"securityKey" gorm:"-"`           // 安全密钥
	Name        string `json:"name"`                            // 银行名称
	RealName    string `json:"realName"`                        // 真实姓名
	Number      string `json:"number"`                          // 证件卡号
	Code        string `json:"code"`                            // 银行代码
	Remark      string `json:"remark"`                          // 备注信息
}

// Update 提现账户更新
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	walletsTemplate := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId).CheckBoxToMaps("walletsTemplate")

	// 是否需要安全密钥
	if walletsTemplate["showAccountSecurityKey"] {
		userInfo := &usersModel.User{}
		database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Find(userInfo)
		if userInfo.SecurityKey != utils.PasswordEncrypt(params.SecurityKey) {
			return ctx.ErrorJsonTranslate("incorrectSecurityKey")
		}
	}

	// 是否可以更新
	if !walletsTemplate["showAccountUpdate"] {
		return ctx.SuccessJsonOK()
	}

	database.Db.Model(&walletsModel.WalletUserAccount{}).Where("id = ?", params.ID).Updates(params)
	return ctx.SuccessJsonOK()
}
