package assets

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/wallets/assets"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: walletsModel.WalletAssetsStatusActive},
	{Label: "禁用", Value: walletsModel.WalletAssetsStatusDisable},
}

var typesOptions = []*views.InputOptions{
	{Label: "数字货币", Value: walletsModel.WalletAssetsTypeDigitalCurrency},
	{Label: "虚拟资产", Value: walletsModel.WalletAssetsTypeVirtualCurrency},
	{Label: "法币资产", Value: walletsModel.WalletAssetsTypeFiatCurrency},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("资产名称", "name").
		Text("资产标识", "symbol").
		Select("资产类型", "type", typesOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增用户资产", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增用户资产").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Image("资产图标", "icon").
					SelectDefault("资产类型", "type", typesOptions).
					Text("资产名称", "name").Text("资产标识", "symbol"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Image("资产图标", "icon", false).
		EditText("资产名称", "name", false).
		EditText("资产标识", "symbol", false).
		EditNumber("资产汇率", "rate", true).
		Select("资产类型", "type", true, typesOptions).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("操作时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Image("资产图标", "icon").Text("资产名称", "name").Text("资产标识", "symbol"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
