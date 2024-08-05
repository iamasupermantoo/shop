package bill

import (
	"gofiber/app/models/model/types"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
	"gorm.io/gorm"
)

type IndexParams struct {
	Mode     int   ` json:"mode" validate:"required"` // 模式
	AssetsId int   `json:"assetsId"`                  // 资产ID
	Types    []int `json:"types"`                     // 类型数组
	context.IndexParams
}

// Index 钱包账单列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*walletUserBill, 0)}
	database.Db.Model(&walletsModel.WalletUserBill{}).Where("user_id = ?", ctx.UserId).
		Preload("AssetsInfo").
		Scopes(func(db *gorm.DB) *gorm.DB {
			if params.Mode == types.WalletsModeBalance {
				return db.Where("assets_id = 0")
			}
			return db.Where("assets_id > 0")
		}).Scopes(scopes.NewScopes().Eq("assets_id = ?", params.AssetsId).
		In("type", params.Types).
		Between("created_at", params.CreatedAt.ToAddDate()).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}

type walletUserBill struct {
	walletsModel.WalletUserBill
	AssetsInfo walletsModel.WalletAssets `json:"assetsInfo" gorm:"foreignKey:ID;references:AssetsId"`
}
