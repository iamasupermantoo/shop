package channel

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` //	管理名称
	Name      string `json:"name"`      //	渠道名称
	Mode      int    `json:"mode"`      //	模式
	Type      int    `json:"type"`      //	类型
	Symbol    string `json:"symbol"`    //	标识
	Status    int    `json:"status"`    //	状态
	context.IndexParams
}

type channel struct {
	usersModel.Channel
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
}

// Index 渠道列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*channel, 0)}
	database.Db.Model(&usersModel.Channel{}).Preload("AdminInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			Eq("name", params.Name).
			Eq("status", params.Status).
			Eq("mode", params.Mode).
			Eq("type", params.Type).
			Eq("symbol", params.Symbol).
			Between("created_at", params.CreatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
