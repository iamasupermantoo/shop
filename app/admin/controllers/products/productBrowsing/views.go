package productBrowsing

import (
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseUrl  = "/auth/products/browsing"
	indexUrl = baseUrl + "/index"
)

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(indexUrl, "")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("店铺名称", "storeName").
		Text("产品名称", "productName").
		RangeDatePicker("更新时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "ID", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("店铺名称", "storeInfo.name", false).
		Text("产品名称", "productInfo.name", false).
		EditNumber("浏览次数", "nums", true).
		DatePicker("更新时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{}...)

	return ctx.SuccessJson(config)
}
