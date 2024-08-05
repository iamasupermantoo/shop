package lang

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理员名称
	Name      string `json:"name"`      // 名称
	Alias     string `json:"alias"`     // 别名
	Sort      int    `json:"sort"`      // 排序
	Status    int    `json:"status"`    // 状态 -1禁用 10开启
	Symbol    string `json:"symbol"`    // 标识
	context.IndexParams
}

type IndexData struct {
	systemsModel.Lang
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}

	database.Db.Model(&systemsModel.Lang{}).
		Preload("AdminInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			Eq("name", params.Name).
			Eq("alias", params.Alias).
			Eq("symbol", params.Symbol).
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
