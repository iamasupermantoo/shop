package translate

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
	Symbol    string `json:"symbol"`    // 语言标识
	Field     string `json:"field"`     // 键名
	Value     string `json:"value"`     // 键值
	Type      int    `json:"type"`      // 类型
	context.IndexParams
}

type IndexData struct {
	systemsModel.Translate
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId" json:"adminInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}
	database.Db.Model(&systemsModel.Translate{}).
		Preload("AdminInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			Eq("name", params.Name).
			Eq("field", params.Field).
			Eq("lang", params.Symbol).
			Eq("type", params.Type).
			Like("value", "%"+params.Value+"%").
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
