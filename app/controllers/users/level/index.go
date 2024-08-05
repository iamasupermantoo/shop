package level

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Index 等级列表
func Index(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	levelList := make([]*systemsModel.SystemLevelInfo, 0)
	database.Db.Model(&systemsModel.Level{}).Where("status=?", systemsModel.LevelStatusActive).
		Where("admin_id=?", ctx.AdminSettingId).Order("symbol ASC").
		Scan(&levelList)
	return ctx.SuccessJson(levelList)
}
