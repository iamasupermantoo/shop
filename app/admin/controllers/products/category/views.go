package category

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/service/productsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/products/category"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: productsModel.CategoryStatusActive},
	{Label: "禁用", Value: productsModel.CategoryStatusDisable},
}

var typeOptions = []*views.InputOptions{
	{Label: "默认分类", Value: productsModel.CategoryTypeDefault},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	categoryOptions := productsService.NewProductCategory(ctx.Rds, ctx.AdminSettingId).GetViewsOptions()

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Select("上级名称", "parentId", categoryOptions).
		Select("类型", "type", typeOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "id", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增分类", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增分类数据").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Image("图标", "icon").
					SelectDefault("父级名称", "parentId", categoryOptions).
					SelectDefault("类型", "type", typeOptions).
					Text("名称(翻译)", "name"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("父级名称", "parentInfo.name", false).
		Image("图标", "icon", false).
		Translate("名称", "name", false).
		EditNumber("排序", "sort", true).
		Select("类型", "type", true, typeOptions).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("时间", "updatedAt", true))

	scopeRowWhere := productsService.NewProductCategory(ctx.Rds, ctx.AdminId).GetEndClientShowWhere(ctx.GetAdminChildIds())

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Image("图标", "icon").
						Select("父级名称", "parentId", categoryOptions).
						Select("类型", "type", typeOptions).
						Text("名称", "name"),
				),
		},
		{
			ButtonViews: views.ButtonViews{Label: "爬取数据", Color: views.ColorSecondary, Size: "xs", Eval: scopeRowWhere},
			Config: views.NewDialogViews("update", baseURL+"/crawling", "爬取分类数据").
				SetSmall("产品地址如：https://www.amazon.com/-/en/dp/B08152D6DL?ref_=Oct_DLandingS_D_948c90c7_0&th=1\nhttps://www.amazon.com/s?k=iphone&__mk_zh_CN=%E4%BA%9A%E9%A9%AC%E9%80%8A%E7%BD%91%E7%AB%99&crid=31MZGQ9HE0U3W&sprefix=%2Caps%2C356&ref=nb_sb_noss_2").
				SetInputViews(
					views.NewInputViews().TextArea("产品Url", "url"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
