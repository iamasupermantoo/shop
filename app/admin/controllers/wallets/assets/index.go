package assets

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` //	管理账户
	Name      string `json:"name"`      // 名称
	Symbol    string `json:"symbol"`    // 标识
	Type      int    `json:"type"`      // 类型 1法币资产 11数字货币资产 21虚拟货币资产
	Status    int    `json:"status"`    // 状态 -1禁用 10开启
	context.IndexParams
}

type walletAssets struct {
	walletsModel.WalletAssets
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*walletAssets, 0)}
	database.Db.Model(&walletsModel.WalletAssets{}).Preload("AdminInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			Eq("name", params.Name).
			Eq("symbol", params.Symbol).
			Eq("type", params.Type).
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
