package payment

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type DataParams struct {
	ID   uint                               `validate:"required" json:"id"` // 支付ID
	Data walletsModel.GormWalletPaymentData `json:"data"`                   // 支付数据
}

func Data(ctx *context.CustomCtx, params *DataParams) error {
	paymentInfo := &walletsModel.WalletPayment{}
	result := database.Db.Model(&walletsModel.WalletPayment{}).Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(paymentInfo)
	if result.Error != nil || paymentInfo.ID == 0 {
		return ctx.ErrorJson("找不到可更新的支付信息")
	}

	result = database.Db.Model(&walletsModel.WalletPayment{}).Where("id = ?", paymentInfo.ID).Update("data", params.Data)
	return ctx.IsErrorJson(result.Error)
}
