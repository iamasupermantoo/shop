package product

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/service/productsService"
	"gofiber/app/models/service/usersService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/products/product"
)

var statusOptions = []*views.InputOptions{
	{Label: "上架", Value: productsModel.ProductStatusActive},
	{Label: "下架", Value: productsModel.ProductStatusDisable},
}

var typeOptions = []*views.InputOptions{
	{Label: "店铺商品", Value: productsModel.ProductTypeDefault},
	{Label: "批发商品", Value: productsModel.ProductTypeWholesale},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	categoryService := productsService.NewProductCategory(ctx.Rds, ctx.AdminSettingId)

	categoryOptions := categoryService.GetViewsOptions()

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("店铺名称", "storeName").
		Text("商品名称", "name").
		Select("分类名称", "categoryId", categoryOptions).
		Select("商品类型", "type", typeOptions).
		Select("商品状态", "status", statusOptions).
		RangeDatePicker("更新时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "新增产品", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增产品数据").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Images("产品图片", "images").
					SelectDefault("分类名称", "categoryId", categoryOptions).
					SelectDefault("类型", "type", typeOptions).
					Text("名称", "name").
					Number("金额", "money"),
			),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "获取产品", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/crawling", "获取产品数据").
				SetSmall("产品地址如：https://www.amazon.com/-/en/dp/B08152D6DL?ref_=Oct_DLandingS_D_948c90c7_0&th=1\nhttps://www.amazon.com/s?k=iphone&__mk_zh_CN=%E4%BA%9A%E9%A9%AC%E9%80%8A%E7%BD%91%E7%AB%99&crid=31MZGQ9HE0U3W&sprefix=%2Caps%2C356&ref=nb_sb_noss_2").
				SetInputViews(
					views.NewInputViews().
						Children("分类", "crawlInfo", views.NewInputViews().
							Select("分类名称", "categoryId", categoryOptions).
							TextArea("产品地址", "urls").GetInputListRows(),
						),
				),
		},
	)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("店铺", "storeInfo.name", false).
		Select("类型", "type", false, typeOptions).
		Select("分类名称", "categoryId", true, categoryOptions).
		Images("产品图片", "images", false).
		EditText("名称", "name", false).
		EditNumber("原价", "money", true).
		EditNumber("折扣(%)", "discount", true).
		EditNumber("销量", "sales", true).
		EditNumber("排序", "sort", true).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("更新时间", "updatedAt", true))

	// 更新产品属性
	attrsValueInput := views.NewInputViews().
		Text("属性名称", "name").GetInputListRows()

	attrsNameInput := views.NewInputViews().
		Text("属性名", "name").
		Children("属性值", "values", attrsValueInput).GetInputListRows()

	// 更新产品SKU
	skuNameInput := views.NewInputViews().
		Images("图片", "images").
		Text("名称", "name").SetReadonly("name").
		Number("库存", "stock").
		Number("原价", "money").
		Number("折扣", "discount").
		Select("状态", "status", statusOptions).GetInputListRows()

	userVirtualOption := usersService.NewUser().GetUserVirtualOption(ctx.AdminId)
	// 数据操作栏目
	config.SetOptions(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "商品信息", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新商品数据信息").
				SetInputViews(
					views.NewInputViews().
						Images("图片组", "images").
						Select("商品分类", "categoryId", categoryOptions).
						Text("商品名称", "name").
						Editor("商品描述", "desc"),
				),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "商品属性", Color: views.ColorSecondary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update/attrs", "更新商品属性信息").SetSizingSmall().
				SetSmall("更新产品属性, 会自动重新生成SKU, 金额自动分配产品金额").
				SetInputViews(
					views.NewInputViews().
						Children("商品属性", "attrs", attrsNameInput),
				),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "商品SKU", Color: views.ColorNegative, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update/sku", "更新商品属性SKU信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Children("商品SKU", "skuList", skuNameInput),
				),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "添加评论", Color: views.ColorPositive, Size: "xs", Eval: "scope.row.storeId > 0"},
			Config: views.NewDialogViews("update", baseURL+"/comment/create", "添加产品评论").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Images("图片", "image").
						Select("虚拟用户", "userId", userVirtualOption).
						Number("评分", "rating").
						DatePicker("更新时间", "createAt").
						TextArea("评论内容", "comment"),
				),
		},
	)

	return ctx.SuccessJson(config)
}
