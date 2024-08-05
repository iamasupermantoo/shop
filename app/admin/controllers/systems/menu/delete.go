package menu

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Delete 删除接口
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	err := database.Db.Where("id IN ?", params.Ids).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).Delete(&systemsModel.Menu{}).Error

	adminsService.NewAdminMenu(ctx.Rds, ctx.AdminSettingId).DelRedisAdminMenuList()
	return ctx.IsErrorJson(err)
}
