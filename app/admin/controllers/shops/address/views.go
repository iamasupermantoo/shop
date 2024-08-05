package address

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL   = "/auth/shops/address"
	indexUrl  = baseURL + "/index"
	updateUrl = baseURL + "/update"
	deleteUrl = baseURL + "/delete"
	createUrl = baseURL + "/delete"
)

var (
	typeOptions = []*views.InputOptions{
		{Label: "收货地址", Value: shopsModel.ShippingAddressTypeReceiving},
		{Label: "发货地址", Value: shopsModel.ShippingAddressTypeShipments},
	}

	statusOptions = []*views.InputOptions{
		{Label: "激活", Value: shopsModel.ShippingAddressStatusActivate},
		{Label: "禁用", Value: shopsModel.ShippingAddressStatusDisabled},
	}

	isShowOptions = []*views.InputOptions{
		{Label: "默认地址", Value: shopsModel.ShippingAddressIsShowYes},
		{Label: "非默认地址", Value: shopsModel.ShippingAddressIsShowNo},
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
		Text("收货/发货名称", "Name").
		Text("联系方式", "contact").
		Text("国家城市", "city").
		Text("详细地址", "address").
		Select("类型", "type", typeOptions).
		Select("状态", "status", statusOptions).
		Select("是否默认", "isShow", isShowOptions).
		RangeDatePicker("更新时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", deleteUrl, "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "Name", Field: "Ids", Scan: "ID"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "用户收件地址", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", createUrl, "新增用户收件地址").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Select("地址类型", "type", typeOptions).SetValue("Type", shopsModel.ShippingAddressTypeReceiving).SetReadonly("type").
					Text("用户账户", "userName").
					Text("收货名称", "name").
					Text("联系方式", "contact").
					Text("国家城市", "city").
					Text("详细地址", "address"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Text("收件人名称", "name", false).
		Text("联系方式", "contact", false).
		Text("国家城市", "city", false).
		Text("详细地址", "address", false).
		Select("类型", "type", false, typeOptions).
		EditToggle("状态", "status", false, statusOptions).
		Select("是否默认", "isShow", false, isShowOptions).
		DatePicker("操作时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", updateUrl, "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Text("收货人姓名", "name").
						Text("联系方式", "contact").
						Text("国家城市", "city").
						Text("详细地址", "address").
						Select("类型", "type", typeOptions).
						Select("状态", "status", statusOptions).
						Select("是否默认", "isShow", isShowOptions),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
