package payment

import (
	"gofiber/app/models/model/types"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/wallets/payment"
)

var paymentModeOptions = []*views.InputOptions{
	{Label: "余额充值模式", Value: walletsModel.WalletPaymentModeDeposit},
	{Label: "资产充值模式", Value: walletsModel.WalletPaymentModeAssetsDeposit},
	{Label: "余额提现模式", Value: walletsModel.WalletPaymentModeWithdraw},
	{Label: "资产提现模式", Value: walletsModel.WalletPaymentModeAssetsWithdraw},
}
var paymentTypeOptions = []*views.InputOptions{
	{Label: "银行卡类型", Value: walletsModel.WalletPaymentTypeBank},
	{Label: "数字货币类型", Value: walletsModel.WalletPaymentTypeDigital},
	{Label: "渠道类型", Value: walletsModel.WalletPaymentTypeChannel},
	{Label: "第三方通道", Value: walletsModel.WalletPaymentTypeThree},
}
var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: walletsModel.WalletPaymentStatusActive},
	{Label: "禁用", Value: walletsModel.WalletPaymentStatusDisable},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	adminChildrenIds := ctx.GetAdminChildIds()
	assetsOptions := walletsService.NewWalletsAssets().AdminAssetsOptions(adminChildrenIds)
	assetsAndDebitCardOptions := append(assetsOptions, &views.InputOptions{Label: "储蓄卡", Value: 0})

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("支付名称", "name").
		Text("支付标识", "symbol").
		Select("资产名称", "assetsId", assetsOptions).
		Select("支付类型", "type", paymentTypeOptions).
		Select("支付模式", "mode", paymentModeOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("操作时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增余额充值方式", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("createBalanceDeposit", baseURL+"/create", "新增余额充值列表").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Image("充值图标", "icon").
					SetValue("mode", walletsModel.WalletPaymentModeDeposit).
					Select("资产名称", "assetsId", assetsAndDebitCardOptions).
					Text("充值名称", "name").
					Text("标识名称", "symbol"),
			),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增资产充值方式", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("createAssetsDeposit", baseURL+"/create", "新增资产充值列表").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().Image("充值图标", "icon").
					SetValue("mode", walletsModel.WalletPaymentModeAssetsDeposit).
					SelectDefault("资产名称", "assetsId", assetsOptions).
					Text("充值名称", "name").
					Text("标识名称", "symbol"),
			),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增余额提现方式", Color: views.ColorSecondary, Size: "md"},
			Config: views.NewDialogViews("createBalanceWithdraw", baseURL+"/create", "新增余额提现列表").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().Image("提现图标", "icon").
					SetValue("mode", walletsModel.WalletPaymentModeWithdraw).
					Select("资产名称", "assetsId", assetsAndDebitCardOptions).
					Text("提现名称", "name").
					Text("标识名称", "symbol"),
			),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增资产提现方式", Color: views.ColorSecondary, Size: "md"},
			Config: views.NewDialogViews("createAssetsWithdraw", baseURL+"/create", "新增资产提现列表").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().Image("提现图标", "icon").
					SetValue("mode", walletsModel.WalletPaymentModeAssetsWithdraw).
					SelectDefault("资产名称", "assetsId", assetsOptions).
					Text("提现名称", "name").
					Text("标识名称", "symbol"),
			),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增渠道充值", Color: views.ColorSecondary, Size: "md"},
			Config: views.NewDialogViews("createChannelDeposit", baseURL+"/create", "新增渠道充值").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().Image("提现图标", "icon").
					SetValue("mode", walletsModel.WalletPaymentModeDeposit).
					SetValue("type", walletsModel.WalletPaymentTypeChannel).
					Text("渠道名称", "name").
					Text("标识名称", "symbol"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Select("支付模式", "mode", true, paymentModeOptions).
		Select("支付类型", "type", true, paymentTypeOptions).
		Select("资产名称", "assetsId", true, assetsAndDebitCardOptions).
		Image("支付图标", "icon", false).
		EditText("支付标识", "symbol", true).
		EditText("支付名称", "name", false).
		EditNumber("汇率", "rate", true).
		EditToggle("显示凭证", "isVoucher", true, []*views.InputOptions{{Label: "显示", Value: types.ModelBoolTrue}, {Label: "隐藏", Value: types.ModelBoolFalse}}).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("操作时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").
				SetInputViews(
					views.NewInputViews().Image("支付图标", "icon").
						Select("支付类型", "type", paymentTypeOptions).
						Select("支付模式", "mode", paymentModeOptions).
						Editor("支付说明", "desc"),
				),
		},
		{
			ButtonViews: views.ButtonViews{Label: "配置", Color: views.ColorSecondary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/data", "更新配置信息").
				SetInputViews(
					views.NewInputViews().
						Children("数据配置", "data", views.NewInputViews().
							Text("名称", "label").SetReadonly("label").
							Text("参数", "field").SetReadonly("field").
							Text("数据", "value").
							Toggle("显示", "isShow", []*views.InputOptions{{Label: "开启", Value: true}, {Label: "关闭", Value: false}}).
							GetInputListColumn()).SetReadonly("data"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
