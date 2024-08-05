package convert

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	context.IndexParams
}

type walletUserConvert struct {
	walletsModel.WalletUserConvert
	AssetsInfo   assetsModel `json:"assetsInfo" gorm:"foreignKey:AssetsId;"`
	ToAssetsInfo assetsModel `json:"toAssetsInfo" gorm:"foreignKey:ToAssetsId;"`
}

type assetsModel struct {
	ID   uint   `json:"id"`   // ID
	Name string `json:"name"` // 资产名称
}

func (assetsModel) TableName() string {
	return "wallet_assets"
}

// Index 资金转换列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*walletUserConvert, 0)}
	database.Db.Model(&walletsModel.WalletUserConvert{}).Where("user_id = ?", ctx.UserId).
		Preload("AssetsInfo").Preload("ToAssetsInfo").
		Scopes(scopes.NewScopes().
			Between("created_at", params.CreatedAt.ToAddDate()).
			Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
