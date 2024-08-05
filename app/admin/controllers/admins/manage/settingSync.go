package manage

import (
	"gofiber/app/models/service/consoleService"
	"gofiber/app/module/context"
)

// SettingSync 商户配置同步
func SettingSync(ctx *context.CustomCtx, params *context.PrimaryKeyParams) error {
	return ctx.IsErrorJson(consoleService.NewMerchant(params.ID, []string{}).RunSync())
}
