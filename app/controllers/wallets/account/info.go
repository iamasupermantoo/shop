package account

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type InfoParams struct {
	ID int `json:"id" validate:"required"` // 提现账户ID
}

// Info 提现账户信息
func Info(ctx *context.CustomCtx, params *InfoParams) error {
	accountInfo := walletsModel.WalletUserAccount{}
	database.Db.Model(&accountInfo).Where("id = ?", params.ID).Where("user_id = ?", ctx.UserId).Find(&accountInfo)

	return ctx.SuccessJson(accountInfo)
}
