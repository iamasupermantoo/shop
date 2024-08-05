package userBill

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理账户
	UserName  string `json:"userName"`  // 用户账户
	AssetsId  int    `json:"assetsId"`  // 资产ID
	Type      int    `json:"type"`      // 账单类型
	context.IndexParams
}

type walletUserBill struct {
	walletsModel.WalletUserBill
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
	UserInfo  usersModel.User       `gorm:"foreignKey:UserId" json:"userInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*walletUserBill, 0)}
	database.Db.Model(&walletsModel.WalletUserBill{}).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Preload("AdminInfo").Preload("UserInfo").
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Eq("type", params.Type).
			Eq("assets_id", params.AssetsId).
			Between("created_at", params.CreatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	for _, item := range data.Items.([]*walletUserBill) {
		item.Name = adminsService.NewAdminSetting(ctx.Rds, item.AdminId).GetRedisAdminSettingField(item.Name)
	}

	return ctx.SuccessJson(data)
}
