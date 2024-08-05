package translate

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Delete 删除接口
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	translateInfo := &systemsModel.Translate{}
	database.Db.Model(translateInfo).Where("id IN ?", params.Ids).Where("admin_id IN ?", ctx.GetAdminChildIds()).Delete(translateInfo)
	// 删除相应的语言包
	systemsService.NewSystemTranslate(ctx.Rds, ctx.AdminSettingId).DelRedisAdminTranslate()
	return ctx.SuccessJsonOK()
}
