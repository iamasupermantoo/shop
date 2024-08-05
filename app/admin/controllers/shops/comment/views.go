package comment

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseUrl   = "/auth/shops/comment"
	indexUrl  = baseUrl + "/index"
	updateUrl = baseUrl + "/update"
	deleteUrl = baseUrl + "/delete"
)

var (
	statusOptions = []*views.InputOptions{
		{Label: "已评论", Value: shopsModel.StoreCommentsStatusComplete},
		{Label: "未评论", Value: shopsModel.StoreCommentsStatusPending},
	}
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
		Text("评论内容", "name").
		Text("订单编号", "orderSn").
		Select("评论状态", "status", statusOptions).
		RangeDatePicker("更新时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", deleteUrl, "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "Name", Field: "Ids", Scan: "ID"}),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Text("店铺名称", "storeInfo.name", false).
		Text("商品名称", "productInfo.name", false).
		Text("订单编号", "orderInfo.orderSn", false).
		EditTextArea("评论内容", "name", false).
		Images("评论图片", "images", false).
		EditNumber("评分", "rating", true).
		Select("状态", "status", false, statusOptions).
		DatePicker("更新时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", updateUrl, "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Images("评论图片", "images").
						TextArea("评论内容", "name"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
