package userLevel

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/users/level"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: usersModel.UserLevelStatusActive},
	{Label: "禁用", Value: usersModel.UserLevelStatusDisable},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	levelOptions := systemsService.NewSystemLevel().GetAdminOptions(ctx.GetAdminChildIds())

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("等级名称", "name").
		Select("等级标识", "symbol", levelOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("购买时间", "createdAt").
		RangeDatePicker("过期时间", "expiredAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增会员等级", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增会员等级").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					SelectDefault("用户等级", "symbol", levelOptions).
					Text("用户账户", "userName"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Image("等级图标", "icon", false).
		Text("等级名称", "name", false).
		Text("涨幅", "Increase", false).
		Select("等级标识", "symbol", false, levelOptions).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("购买时间", "createdAt", true).
		DatePicker("过期时间", "expiredAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Select("等级标识", "symbol", levelOptions).
						Number("涨幅", "Increase").
						DatePicker("过期时间", "expiredAt"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
