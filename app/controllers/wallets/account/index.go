package account

import (
	"gofiber/app/models/model/types"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
)

type IndexParams struct {
	Mode int `json:"mode" validate:"required"` // 提现模式
}

// Index 提现账户列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	paymentIds := make([]uint, 0)
	optionsList := make([]*views.InputOptions, 0)
	walletsInstance := walletsService.NewWalletsAssets()
	if params.Mode == types.WalletsModeBalance {
		optionsList = walletsInstance.GetBalanceWithdrawOptions([]uint{ctx.AdminSettingId})
	} else {
		optionsList = walletsInstance.GetAssetsWithdrawOptions([]uint{ctx.AdminSettingId})
	}
	for _, options := range optionsList {
		paymentIds = append(paymentIds, options.Value.(uint))
	}

	data := make([]*userAccount, 0)
	database.Db.Model(&walletsModel.WalletUserAccount{}).Preload("PaymentInfo").
		Where("status = ?", walletsModel.WalletUserAccountStatusActive).
		Where("payment_id IN ?", paymentIds).
		Where("user_id = ?", ctx.UserId).Find(&data)

	return ctx.SuccessJson(data)
}

type userAccount struct {
	walletsModel.WalletUserAccount
	PaymentInfo walletsModel.WalletPayment `json:"paymentInfo" gorm:"foreignKey:PaymentId"`
}
