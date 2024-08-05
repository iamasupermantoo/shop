package article

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理账户
	Type      int    `json:"type"`      // 1基础文章
	Status    int    `json:"status"`    // 状态 -1禁用 10开启
	context.IndexParams
}

type article struct {
	systemsModel.Article
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*article, 0)}
	database.Db.Model(&systemsModel.Article{}).Preload("AdminInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			Eq("type", params.Type).
			Eq("status", params.Status).
			Between("created_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
