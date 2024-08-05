package menu

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Delete 删除接口
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	database.Db.Where("id IN ?", params.Ids).Delete(&adminsModel.AdminMenu{})

	// 删除所有管理菜单
	adminList := make([]*adminsModel.AdminUser, 0)
	database.Db.Model(&adminsModel.AdminUser{}).Find(&adminList)
	for _, adminInfo := range adminList {
		adminsService.NewAdminMenu(ctx.Rds, adminInfo.ID).DelRedisAdminMenuList()
	}
	return ctx.SuccessJsonOK()
}
