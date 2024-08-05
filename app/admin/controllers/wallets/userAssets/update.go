package userAssets

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type UpdateParams struct {
	ID     uint `gorm:"-" validate:"required" json:"id"` //	ID
	Status int  `json:"status"`                          // 状态
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&walletsModel.WalletUserAssets{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
