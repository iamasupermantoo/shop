package order

import (
	"gofiber/app/models/model/types"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AssetsId   int                `json:"assetsId"`   // 资产ID
	Mode       int                `json:"mode"`       // 1 = 1充值类型 11提现类型 | 2 =  2资产充值类型 12资产提现类型
	Pagination *scopes.Pagination `json:"pagination"` // 分页数据
}

// Index 钱包订单列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*walletsModel.WalletUserOrderInfo, 0)}

	typeMode := []int{walletsModel.WalletUserOrderTypeDeposit, walletsModel.WalletUserOrderTypeWithdraw}
	if params.Mode == types.WalletsModeAssets {
		typeMode = []int{walletsModel.WalletUserOrderTypeAssetsDeposit, walletsModel.WalletUserOrderTypeAssetsWithdraw}
	}

	database.Db.Model(&walletsModel.WalletUserOrder{}).Preload("PaymentInfo").Preload("AccountInfo").
		Where("assets_id = ?", params.AssetsId).
		Where("type IN ?", typeMode).Where("user_id = ?", ctx.UserId).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	for _, order := range data.Items.([]*walletsModel.WalletUserOrderInfo) {
		switch order.Type {
		case walletsModel.WalletUserOrderTypeWithdraw:
			database.Db.Model(&walletsModel.WalletPayment{}).Where("id = ?", order.AccountInfo.PaymentId).Find(&order.PaymentInfo)
		}
	}

	return ctx.SuccessJson(data)
}
