package index

import (
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
)

// Download 下载
func Download(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	data := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId).GetDownload()
	return ctx.SuccessJson(data)
}
