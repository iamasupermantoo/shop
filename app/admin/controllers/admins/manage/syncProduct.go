package manage

import (
	"go.uber.org/zap"
	"gofiber/app/models/service/consoleService"
	"gofiber/app/models/service/shopsService"
	"gofiber/app/module/context"
)

// InitProduct 同步产品
func InitProduct(ctx *context.CustomCtx, params *context.PrimaryKeyParams) error {
	go func() {
		err := consoleService.NewMerchant(params.ID, []string{"admin_setting", "wallet_assets", "wallet_payment", "translate", "lang", "country", "level", "article", "menu"}).RunRest()
		if err != nil {
			return
		}

		err = shopsService.NewStoresService(nil, params.ID).DropStore().InitStore()
		if err != nil {
			zap.L().Error("initProduct", zap.Error(err))
			return
		}
	}()
	return ctx.SuccessJsonOK()
}
