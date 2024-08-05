package userConvert

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName  string `json:"adminName"`  // 管理账户
	UserName   string `json:"userName"`   // 用户账户
	AssetsId   int    `json:"assetsId"`   // 资产ID
	ToAssetsId int    `json:"toAssetsId"` // 接收资产ID
	Status     int    `json:"status"`     // 状态 -1失败 10完成
	context.IndexParams
}

type walletUserConvert struct {
	walletsModel.WalletUserConvert
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
	UserInfo  usersModel.User       `gorm:"foreignKey:UserId" json:"userInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*walletUserConvert, 0)}
	database.Db.Model(&walletsModel.WalletUserConvert{}).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Preload("AdminInfo").Preload("UserInfo").
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Eq("assets_id", params.AssetsId).
			Eq("to_assets_id", params.ToAssetsId).
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
