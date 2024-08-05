package refund

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseUrl   = "/auth/shops/refund"
	indexUrl  = baseUrl + "/index"
	updateUrl = baseUrl + "/update"
	deleteUrl = baseUrl + "/delete"
)

var (
	statusOptions = []*views.InputOptions{
		{Label: "申请", Value: shopsModel.StoreRefundStatusPending},
		{Label: "拒绝", Value: shopsModel.StoreRefundStatusRefuse},
		{Label: "同意", Value: shopsModel.StoreRefundStatusComplete}}
)

// Views 视图配置
func Views(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(indexUrl, updateUrl)

	// 查询设置
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("店铺名称", "storeName").
		Text("订单编号", "orderSn").
		Text("申请理由", "name").
		Select("售后状态", "status", statusOptions).
		RangeDatePicker("更新时间", "updatedAt").
		GetInputListInfo()

	// 头部操作按钮
	config.Table.Tools = []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", deleteUrl, "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "id", Field: "Ids", Scan: "id"}),
		},
	}

	// 数据表格
	config.Table.Columns = views.NewColumnsViews().
		Text("主键", "id", false).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Text("店铺名称", "storeInfo.name", false).
		Text("订单编号", "orderInfo.orderSn", true).
		EditTextArea("申请理由", "name", false).
		Images("申请凭证", "images", false).
		Text("退款金额", "money", false).
		Select("售后状态", "status", false, statusOptions).
		DatePicker("更新时间", "updatedAt", true).
		GetColumnsListInfo()

	// 数据操作栏目
	config.Table.Options = []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", updateUrl, "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Images("申请凭证", "images").
						TextArea("申请理由", "name").
						Select("状态", "status", statusOptions),
				),
		},
	}

	return ctx.SuccessJson(config)
}
