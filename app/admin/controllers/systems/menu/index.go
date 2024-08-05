package menu

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理账户
	ParentId  int    `json:"parentId"`  // 父级ID
	Route     string `json:"route"`     // 路由
	Type      int    `json:"type"`      // 类型1导航菜单 11设置菜单 21更多菜单
	Status    int    `json:"status"`    // 状态-1禁用 10开启
	context.IndexParams
}

type menu struct {
	systemsModel.Menu
	AdminInfo  adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
	ParentInfo *systemsModel.Menu    `gorm:"foreignKey:ID;references:ParentId" json:"parentInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*menu, 0)}
	database.Db.Model(&systemsModel.Menu{}).Preload("AdminInfo").Preload("ParentInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			Eq("parent_id", params.ParentId).
			Eq("route", params.Route).
			Eq("type", params.Type).
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
