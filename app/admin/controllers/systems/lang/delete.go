package lang

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Delete 删除接口
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	database.Db.Where("id IN ?", params.Ids).Where("admin_id IN ?", ctx.GetAdminChildIds()).Delete(&systemsModel.Lang{})
	systemsService.NewSystemLang(ctx.Rds, ctx.AdminSettingId).DelRedisAdminLangList()

	return ctx.SuccessJsonOK()
}
