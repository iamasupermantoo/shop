package channel

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/users/channel"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: usersModel.ChannelStatusActive},
	{Label: "禁用", Value: usersModel.ChannelStatusDisable},
}
var typeOptions = []*views.InputOptions{
	{Label: "默认类型", Value: usersModel.ChannelTypeDefault},
}
var modeOptions = []*views.InputOptions{
	{Label: "授权", Value: usersModel.ChannelModeApprove},
	{Label: "渠道", Value: usersModel.ChannelModeChannel},
}

func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("渠道名称", "name").
		Text("渠道标识", "symbol").
		Select("模式", "mode", modeOptions).
		Select("类型", "type", typeOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("创建时间", "createdAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增授权", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("createApprove", baseURL+"/create", "新增授权数据").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Text("授权名称", "name").Text("授权标识", "symbol").
					Text("授权路由", "route").SetValue("mode", usersModel.ChannelModeApprove),
			),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增渠道", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("createChannel", baseURL+"/create", "新增渠道数据").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Text("渠道名称", "name").Text("渠道标识", "symbol").
					Text("渠道路由", "route").SetValue("mode", usersModel.ChannelModeChannel),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Select("类型", "type", true, typeOptions).
		Select("模式", "mode", true, modeOptions).
		EditText("名称", "name", false).
		EditText("标识", "symbol", false).
		EditText("链接", "route", false).
		EditText("密码", "pass", false).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("创建时间", "createdAt", true))

	return ctx.SuccessJson(config)
}
