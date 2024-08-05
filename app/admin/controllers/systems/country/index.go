package country

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

// IndexParams 国家列表参数
type IndexParams struct {
	AdminName string `json:"adminName"` // 管理员名称
	Name      string `json:"name"`      // 名称
	Alias     string `json:"alias"`     // 别名
	Iso1      string `json:"iso1"`      // 二位代码
	Code      string `json:"code"`      // 区号
	Status    int    `json:"status"`    // 状态
	context.IndexParams
}

type IndexData struct {
	systemsModel.Country
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}
	database.Db.Model(&systemsModel.Country{}).
		Preload("AdminInfo").
		Where("admin_id in ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			Eq("name", params.Name).
			Eq("alias", params.Alias).
			Eq("code", params.Code).
			Eq("iso1", params.Iso1).
			Eq("status", params.Status).
			Between("created_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
