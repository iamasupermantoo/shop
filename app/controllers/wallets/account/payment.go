package account

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type PaymentParams struct {
	Mode int `json:"mode" validate:"required"` // 模式 1充值模式 2资产充值模式 11提现模式 12资产提现模式
}

// Payment 提现账户类型
func Payment(ctx *context.CustomCtx, params *PaymentParams) error {
	data := make([]*walletsModel.WalletPaymentInfo, 0)
	database.Db.Model(&walletsModel.WalletPayment{}).Where("mode = ?", params.Mode).
		Where("admin_id = ?", ctx.AdminSettingId).Where("status = ?", walletsModel.WalletPaymentStatusActive).Find(&data)
	return ctx.SuccessJson(data)
}
