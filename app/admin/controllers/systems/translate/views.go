package translate

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/systems/translate"
)

var typeOptions = []*views.InputOptions{
	{Label: "默认翻译", Value: systemsModel.TranslateTypeDefault},
	{Label: "前台翻译", Value: systemsModel.TranslateTypeFrontend},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	langOptions := systemsService.NewSystemLang(ctx.Rds, ctx.AdminSettingId).GetAdminOptions(ctx.GetAdminChildIds())

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("名称", "name").
		Text("键名", "field").
		SetValue("field", ctx.Query("field")).
		Text("键值", "value").
		Select("类型", "type", typeOptions).
		Select("标识", "symbol", langOptions).
		RangeDatePicker("时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "id", Field: "ids", Scan: "id"}),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "新增数据", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增数据").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					SelectDefault("类型", "type", typeOptions).
					SelectDefault("标识", "lang", langOptions).
					Text("名称", "name").
					Text("键名", "field").
					Text("键值", "value"),
			),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "新增语言翻译", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("createLang", baseURL+"/lang", "新增语言翻译").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Select("标识", "lang", langOptions),
			),
		},
	)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		EditText("名称", "name", false).
		Select("类型", "type", true, typeOptions).
		Text("标识", "lang", false).
		Text("键名", "field", false).
		EditTextArea("键值", "value", false).
		DatePicker("时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Select("类型", "type", typeOptions).
						Select("标识", "lang", langOptions).
						Text("名称", "name").
						Editor("键值", "value"),
				),
		},
	)

	return ctx.SuccessJson(config)
}
