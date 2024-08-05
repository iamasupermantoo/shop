package article

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/systems/article"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: systemsModel.ArticleStatusActive},
	{Label: "禁用", Value: systemsModel.ArticleStatusDisable},
}

var typeOptions = []*views.InputOptions{
	{Label: "基础文章", Value: systemsModel.ArticleTypeDefault},
	{Label: "帮助", Value: systemsModel.ArticleTypeHelpers},
	{Label: "关于", Value: systemsModel.ArticleTypeAbout},
	{Label: "产品", Value: systemsModel.ArticleTypeProduct},
	{Label: "服务", Value: systemsModel.ArticleTypeService},
	{Label: "社交", Value: systemsModel.ArticleTypeSocial},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Select("类型", "type", typeOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "id", Field: "ids", Scan: "id"}),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "新增文章", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增文章").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Image("封面", "image").
					SelectDefault("类型", "type", typeOptions).
					Text("标题(翻译)", "name").
					Text("内容(翻译)", "content"),
			),
		},
	)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Select("类型", "type", true, typeOptions).
		Image("封面", "image", false).
		Translate("标题(翻译)", "name", false).
		Translate("内容(翻译)", "content", false).
		EditText("链接", "link", false).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().Image("封面", "image").
						Text("标题", "name").
						Select("类型", "type", typeOptions),
				),
		},
	)

	return ctx.SuccessJson(config)
}
