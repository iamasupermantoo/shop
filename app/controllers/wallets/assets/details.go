package assets

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type DetailsParams struct {
	ID int `json:"id" validate:"required" gorm:"-"` // 资产ID
}

// Details 用户资产详情
func Details(ctx *context.CustomCtx, params *DetailsParams) error {
	data := &assetsInfo{}
	database.Db.Model(&walletsModel.WalletAssets{}).Where("id = ?", params.ID).Where("admin_id = ?", ctx.AdminSettingId).
		Preload("UserAssets", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", ctx.UserId)
		}).Find(data)

	return ctx.SuccessJson(data)
}
