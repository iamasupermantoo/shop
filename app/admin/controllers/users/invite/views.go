package invite

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/users/invite"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: usersModel.InviteStatusActive},
	{Label: "禁用", Value: usersModel.InviteStatusDisable},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("邀请码", "code").
		Select("状态", "status", statusOptions).
		RangeDatePicker("操作时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "code", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增管理邀请码", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("createAdmin", baseURL+"/create", "新增管理邀请码").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().Text("管理账户", "adminName"),
			),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增用户邀请码", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("createUser", baseURL+"/create", "新增用户邀请码").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().Text("用户账户", "userName"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理用户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		EditText("邀请码", "code", false).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("操作时间", "updatedAt", true))

	return ctx.SuccessJson(config)
}
