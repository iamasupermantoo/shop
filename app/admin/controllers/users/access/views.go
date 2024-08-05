package access

import (
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/users/access"
)

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("记录名称", "name").
		Text("IP地址", "ip").
		Text("记录路由", "route").
		RangeDatePicker("记录时间", "createdAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Text("记录名称", "name", false).
		Text("IP地址", "ip", true).
		Text("记录路由", "route", false).
		Text("Headers", "headers", false).
		DatePicker("记录时间", "createdAt", true))

	return ctx.SuccessJson(config)
}
