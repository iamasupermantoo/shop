package manage

import (
	"gofiber/app/models/service/consoleService"
	"gofiber/app/module/context"
)

// SettingReset 商户配置重置
func SettingReset(ctx *context.CustomCtx, params *context.PrimaryKeyParams) error {
	return ctx.IsErrorJson(consoleService.NewMerchant(params.ID, []string{}).RunRest())
}
