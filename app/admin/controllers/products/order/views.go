package order

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/products/order"
)

var typeOptions = []*views.InputOptions{
	{Label: "商家订单", Value: productsModel.ProductOrderTypeDefault},
}

var statusOptions = []*views.InputOptions{
	{Label: "已取消", Value: productsModel.ProductOrderStatusDisable},
	{Label: "待付款", Value: productsModel.ProductOrderStatusPending},
	{Label: "待发货", Value: productsModel.ProductOrderStatusShipping},
	{Label: "待收货", Value: productsModel.ProductOrderStatusProgress},
	{Label: "已完成", Value: productsModel.ProductOrderStatusComplete},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("订单编号", "orderSn").
		Number("产品ID", "productId").
		Select("订单类型", "type", typeOptions).
		Select("订单状态", "status", statusOptions).
		RangeDatePicker("购买时间", "createdAt"))

	// 头部操作按钮
	config.SetTools(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "orderSn", Field: "ids", Scan: "id"}),
		},
		//&views.DialogButtonViews{
		//	ButtonViews: views.ButtonViews{Label: "新增订单", Color: views.ColorPrimary, Size: "md"},
		//	Config: views.NewDialogViews("create", baseURL+"/create", "新增订单数据").
		//		SetSizingSmall().SetInputViews(
		//		views.NewInputViews().
		//			SelectDefault("订单类型", "type", typeOptions).
		//			Text("用户账户", "userName").
		//			Number("产品ID", "productId"),
		//	),
		//},
	)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Images("产品图片", "productInfo.images", false).
		Text("用户账户", "userInfo.userName", false).
		Text("订单编号", "orderSn", false).
		Text("名称", "productInfo.name", false).
		EditNumber("金额", "money", true).
		Select("类型", "type", true, typeOptions).
		Select("状态", "status", true, statusOptions).
		DatePicker("购买时间", "createdAt", true).
		DatePicker("付款时间", "expiredAt", true))

	// 数据操作栏目
	config.Table.Options = []*views.DialogButtonViews{}

	return ctx.SuccessJson(config)
}
