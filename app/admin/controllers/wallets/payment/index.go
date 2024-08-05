package payment

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理账户
	AssetsId  int    `json:"assetsId"`  // 资产名称
	Name      string `json:"name"`      // 名称
	Symbol    string `json:"symbol"`    // 标识
	Type      int    `json:"type"`      // 类型 1银行卡类型 11数字货币类型 21第三方支付
	Mode      int    `json:"mode"`      // 模式 1充值模式 2资产充值模式 11提现模式 12资产提现模式
	Status    int    `json:"status"`    // 状态 -1禁用 10开启
	context.IndexParams
}

type walletPayment struct {
	walletsModel.WalletPayment
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*walletPayment, 0)}
	database.Db.Model(&walletsModel.WalletPayment{}).Preload("AdminInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			Eq("assets_id", params.AssetsId).
			Eq("name", params.Name).
			Eq("symbol", params.Symbol).
			Eq("type", params.Type).
			Eq("mode", params.Mode).
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
