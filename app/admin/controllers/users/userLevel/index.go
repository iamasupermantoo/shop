package userLevel

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string                  `json:"adminName"` // 管理账户
	UserName  string                  `json:"userName"`  // 用户账户
	Name      string                  `json:"name"`      // 等级名称
	Symbol    int                     `json:"symbol"`    // 标识
	Status    int                     `json:"status"`    // 状态 -1禁用 10开启
	ExpiredAt *scopes.RangeDatePicker `json:"expiredAt"` // 过期时间
	context.IndexParams
}

type userLevel struct {
	usersModel.UserLevel
	AdminInfo *adminsModel.AdminUser `gorm:"foreignKey:AdminId" json:"adminInfo"`
	UserInfo  *usersModel.User       `gorm:"foreignKey:UserId" json:"userInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*userLevel, 0)}
	database.Db.Model(&usersModel.UserLevel{}).Preload("AdminInfo").Preload("UserInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Eq("name", params.Name).
			Eq("symbol", params.Symbol).
			Eq("status", params.Status).
			Between("created_at", params.CreatedAt).
			Between("expired_at", params.ExpiredAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
