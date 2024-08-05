package setting

import (
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/admins/setting"
)

// Views 配置视图
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	adminSettingService := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId)
	data := &viewsData{
		IndexURL:     baseURL + "/index",
		CreateURL:    baseURL + "/create",
		UpdateURL:    baseURL + "/update",
		DeleteURL:    baseURL + "/delete",
		GroupOptions: adminSettingService.GroupOptions(),
	}

	return ctx.SuccessJson(data)
}

type viewsData struct {
	IndexURL     string //	请求路由
	CreateURL    string //	创建路由
	UpdateURL    string //	更新路由
	DeleteURL    string //	删除路由
	GroupOptions []*views.InputOptions
}
