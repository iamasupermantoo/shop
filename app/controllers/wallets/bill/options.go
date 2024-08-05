package bill

import (
	"gofiber/app/models/model/types"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
)

type OptionsParams struct {
	Mode int `json:"mode" validate:"required"`
}

// Options 钱包账单类型
func Options(ctx *context.CustomCtx, params *OptionsParams) error {
	userBillCache := walletsService.NewUserBill()
	if params.Mode == types.WalletsModeBalance {
		return ctx.SuccessJson(userBillCache.GetBalanceOptions())
	}

	return ctx.SuccessJson(userBillCache.GetAssetsOptions())
}
