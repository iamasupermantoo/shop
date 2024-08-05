package systems

import (
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
)

// Translate 获取语言包
func Translate(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	langList := systemsService.NewSystemTranslate(ctx.Rds, ctx.AdminSettingId).GetRedisAdminTranslateLangList(ctx.Lang)
	return ctx.SuccessJson(langList)
}
