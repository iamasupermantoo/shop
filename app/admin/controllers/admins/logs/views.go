package logs

import (
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/admins/logs"
)

// Views 视图配置
func Views(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	config.SetSearch(views.NewInputViews().
		Text("管理账户", "userName").
		Text("日志标题", "name").
		Text("IP地址", "ip4").
		Text("日志路由", "route").
		Text("请求信息", "headers").
		RangeDatePicker("时间", "createdAt"))

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("日志标题", "name", false).
		Text("日志路由", "route", false).
		Text("IP地址", "ip", false).
		Text("请求信息", "headers", false).
		Text("请求参数", "body", false).
		DatePicker("时间", "createdAt", true))

	// 数据操作栏目
	config.Table.Options = []*views.DialogButtonViews{}

	return ctx.SuccessJson(config)
}
