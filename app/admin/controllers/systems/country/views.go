package country

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/systems/country"
)

var statusOptions = []*views.InputOptions{
	{Label: "开启", Value: systemsModel.CountryStatusActive},
	{Label: "禁用", Value: systemsModel.CountryStatusDisable},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")
	config.Pagination = &views.Pagination{SortBy: "status", Descending: true, Page: 1, RowsPerPage: 20}

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("国家名称", "name").
		Text("国家别名", "alias").
		Text("二位代码", "iso1").
		Text("区号", "code").
		Select("状态", "status", statusOptions).
		RangeDatePicker("操作时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增数据", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增数据").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Image("国家图标", "icon").
					Text("国家名称", "name").
					Text("国家别名", "alias").
					Text("二位代码", "iso1").
					Text("国家区号", "code"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Image("图标", "icon", false).
		EditText("国家名字", "name", true).
		EditText("国家别名", "alias", true).
		EditText("二位代码", "iso1", true).
		EditNumber("排序", "sort", true).
		EditText("区号", "code", true).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("操作时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Image("国家图标", "icon").
						Text("国家名字", "name").
						Text("国家别名", "alias"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
