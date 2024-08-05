package assets

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

// Assets 钱包资产列表
func Assets(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	data := make([]*assetsInfo, 0)
	database.Db.Model(&walletsModel.WalletAssets{}).Where("admin_id = ?", ctx.AdminSettingId).
		Preload("UserAssets", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", ctx.UserId)
		}).Where("status = ?", walletsModel.WalletAssetsStatusActive).Find(&data)

	return ctx.SuccessJson(data)
}

type assetsInfo struct {
	walletsModel.WalletAssets
	UserAssets walletsModel.WalletUserAssets `json:"userAssets" gorm:"foreignKey:AssetsId"`
}

func (assetsInfo) TableName() string {
	return "assets"
}
