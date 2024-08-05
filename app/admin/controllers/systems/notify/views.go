package notify

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/systems/notify"
)

var statusOptions = []*views.InputOptions{
	{Label: "未读消息", Value: systemsModel.NotifyStatusActive},
	{Label: "已读消息", Value: systemsModel.NotifyStatusComplete},
}

var modeOptions = []*views.InputOptions{
	{Label: "前台", Value: systemsModel.NotifyModeHomeMessage},
	{Label: "后台", Value: systemsModel.NotifyModeAdminMessage},
}
var frontModeOptions = []*views.InputOptions{{Label: "前台", Value: systemsModel.NotifyModeHomeMessage}}
var backModeOptions = []*views.InputOptions{{Label: "后台", Value: systemsModel.NotifyModeAdminMessage}}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("通知标题", "name").
		Select("通知模式", "mode", modeOptions).
		Select("通知状态", "status", statusOptions).
		RangeDatePicker("发送时间", "createdAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增管理通知", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增管理通知").
				SetInputViews(
					views.NewInputViews().
						Select("通知模式", "mode", backModeOptions).
						SetValue("mode", systemsModel.NotifyModeAdminMessage).
						Text("用户账户", "username").
						Text("通知标题", "name").
						Editor("通知内容", "content"),
				),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增用户通知", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增用户通知").
				SetInputViews(
					views.NewInputViews().
						Select("通知模式", "mode", frontModeOptions).
						SetValue("mode", systemsModel.NotifyModeHomeMessage).
						Text("用户账户", "username").
						Text("通知标题", "name").
						Editor("通知内容", "content"),
				),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		EditText("通知标题", "name", false).
		Select("通知状态", "status", true, statusOptions).
		Select("通知模式", "mode", false, modeOptions).
		EditTextArea("通知内容", "content", false).
		DatePicker("发送时间", "createdAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").
				SetInputViews(
					views.NewInputViews().
						Select("通知状态", "status", statusOptions).
						Editor("通知内容", "content"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
