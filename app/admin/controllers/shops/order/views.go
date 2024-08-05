package shopsOrder

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseUrl     = "/auth/shops/order"
	indexUrl    = baseUrl + "/index"
	shippingUrl = baseUrl + "/shipping"
	completeUrl = baseUrl + "/complete"
)

var (
	statusOptions = []*views.InputOptions{
		{Label: "订单取消", Value: productsModel.ProductOrderStatusDisable},
		{Label: "待付款", Value: productsModel.ProductOrderStatusPending},
		{Label: "待发货", Value: productsModel.ProductOrderStatusShipping},
		{Label: "待收货", Value: productsModel.ProductOrderStatusProgress},
		{Label: "订单完成", Value: productsModel.ProductOrderStatusComplete},
	}
)

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(indexUrl, "")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("店铺名称", "storeName").
		Text("店铺订单编号", "orderSn").
		Select("商品状态", "status", statusOptions).
		RangeDatePicker("更新时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Text("店铺名称", "storeInfo.name", false).
		Text("店铺订单编号", "orderSn", false).
		Text("购买总价", "money", true).
		Select("状态", "status", true, statusOptions).
		DatePicker("更新时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions(&views.DialogButtonViews{
		ButtonViews: views.ButtonViews{Label: "发货", Color: views.ColorPrimary, Size: "xs"},
		Config:      views.NewDialogViews("shipping", shippingUrl, "开始发货").SetSizingSmall(),
	},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "收货", Color: views.ColorSecondary, Size: "xs"},
			Config:      views.NewDialogViews("complete", completeUrl, "确认收货").SetSizingSmall(),
		})

	return ctx.SuccessJson(config)
}
