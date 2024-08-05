package cart

import (
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseUrl   = "auth/shops/cart"
	indexUrl  = baseUrl + "/index"
	updateUrl = baseUrl + "/update"
	deleteUrl = baseUrl + "/delete"
)

// Views 视图配置
func Views(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(indexUrl, updateUrl)

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("店铺名称", "storeName").
		Text("商品名称", "productName").
		RangeDatePicker("更新时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除购物车", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", deleteUrl, "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "ID", Field: "Ids", Scan: "ID"}),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Text("店铺名称", "storeInfo.name", false).
		Images("商品图片", "productInfo.images", false).
		Text("商品名称", "productInfo.name", false).
		Text("规格名称", "skuInfo.Name", false).
		EditNumber("规格数量", "nums", false).
		DatePicker("更新时间", "updatedAt", true))

	return ctx.SuccessJson(config)
}
