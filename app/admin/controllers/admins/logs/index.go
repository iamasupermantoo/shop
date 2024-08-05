package logs

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	UserName string `json:"userName"` // 管理名称
	Name     string `json:"name"`     // 管理标题
	Ip4      string `json:"ip4"`      // IP4
	Headers  string `json:"headers"`  // 头信息
	Route    string `json:"route"`    // 路由
	context.IndexParams
}

type adminLogs struct {
	adminsModel.AdminLogs
	AdminInfo *adminsModel.AdminUser `gorm:"foreignKey:AdminId" json:"adminInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*adminLogs, 0)}
	//	过滤参数
	database.Db.Model(&adminsModel.AdminLogs{}).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Preload("AdminInfo").
		Scopes(scopes.NewScopes().FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name", params.UserName).
			Eq("name", params.Name).
			Eq("ip", params.Ip4).
			Like("headers", "%"+params.Headers+"%").
			Eq("route", params.Route).
			Between("created_at", params.CreatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
