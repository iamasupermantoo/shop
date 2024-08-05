package manage

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/consoleService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Delete 删除管理
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	for _, id := range params.Ids {
		database.Db.Where("id = ?", id).Where("id IN ?", ctx.GetAdminChildIds()).Delete(&adminsModel.AdminUser{})
		// 当前客户是超级管理员
		if ctx.AdminId == adminsModel.SuperAdminId {
			_ = consoleService.NewMerchant(uint(id), []string{}).Delete()
		}
	}

	return ctx.SuccessJsonOK()
}
