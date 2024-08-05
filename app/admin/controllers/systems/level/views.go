package level

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/systems/level"
)

var statusOptions = []*views.InputOptions{
	{Label: "开启", Value: systemsModel.LevelStatusActive},
	{Label: "禁用", Value: systemsModel.LevelStatusDisable},
}

var typeOptions = []*views.InputOptions{
	{Label: "全额购买", Value: systemsModel.LevelTypeFullPrice},
	{Label: "等额购买", Value: systemsModel.LevelTypeDifference},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("名称", "name").
		Select("购买方式", "type", typeOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("操作时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "新增数据", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增数据").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Image("图标", "icon").
					Text("名称", "name").
					Number("金额", "money"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Image("图标", "icon", false).
		EditNumber("标识", "symbol", true).
		EditText("名称", "name", true).
		EditNumber("金额", "money", true).
		Text("涨幅", "Increase", false).
		EditNumber("天数", "days", true).
		Select("购买方式", "type", true, typeOptions).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("操作时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").
				SetInputViews(
					views.NewInputViews().
						Image("图标", "icon").
						Select("购买方式", "type", typeOptions).
						Number("涨幅", "Increase").
						Editor("详情", "desc"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
