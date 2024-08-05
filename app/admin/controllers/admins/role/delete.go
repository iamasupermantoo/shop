package role

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Delete 删除接口
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	for _, id := range params.Ids {
		authChildInfo := &adminsModel.AuthChild{}
		result := database.Db.Model(authChildInfo).Where("id = ?", id).Find(authChildInfo)
		if result.Error == nil && result.RowsAffected > 0 {
			database.Db.Where("id = ?", authChildInfo.ID).Delete(&adminsModel.AuthChild{})

			// 清除角色缓存的 - 查询当前角色有哪些管理
			roleAdminList := make([]*adminsModel.AuthAssignment, 0)
			database.Db.Model(&adminsModel.AuthAssignment{}).Where("name = ?", authChildInfo.Child).Find(&roleAdminList)
			for _, assignment := range roleAdminList {
				adminsService.NewAdminMenu(ctx.Rds, assignment.AdminId).DelRedisAdminMenuList()
			}
		}
	}

	return ctx.SuccessJsonOK()
}
