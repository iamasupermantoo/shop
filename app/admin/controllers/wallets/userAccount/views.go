package userAccount

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/users/account"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: walletsModel.WalletPaymentStatusActive},
	{Label: "禁用", Value: walletsModel.WalletPaymentStatusDisable},
}

var modeOptions = []*views.InputOptions{
	{Label: "余额提现", Value: walletsModel.WalletPaymentModeWithdraw},
	{Label: "资产提现", Value: walletsModel.WalletPaymentModeAssetsWithdraw},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	walletService := walletsService.NewWalletsAssets()
	adminChildIds := ctx.GetAdminChildIds()
	balanceOptions := walletService.GetBalanceWithdrawOptions(adminChildIds)
	assetsOptions := walletService.GetAssetsWithdrawOptions(adminChildIds)
	withdrawOptions := append(balanceOptions, assetsOptions...)

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Select("提现类型", "mode", modeOptions).
		Select("提现名称", "paymentId", withdrawOptions).
		Text("账户名称", "name").
		Text("账户姓名", "realName").
		Text("卡号|地址", "number").
		Text("账户代码", "code").
		Text("账户备注", "remark").
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
			ButtonViews: views.ButtonViews{Label: "新增余额提现账户", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("createBalance", baseURL+"/create", "新增余额提现账户").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					SelectDefault("提现名称", "paymentId", balanceOptions).
					Text("用户账户", "userName").
					Text("账户名称", "name").
					Text("账户姓名", "realName").
					Text("卡号｜地址", "number").
					Text("账户代码", "code").
					Text("账户备注", "remark"),
			),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增资产提现账户", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("createAssets", baseURL+"/create", "新增资产提现账户").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Select("提现名称", "paymentId", assetsOptions).SetValue("paymentId", assetsOptions[0].Value).
					Text("用户账户", "userName").
					Text("账户名称", "name").
					Text("账户姓名", "realName").
					Text("卡号｜地址", "number").
					Text("账户代码", "code").
					Text("账户备注", "remark"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Select("提现类型", "paymentInfo.mode", false, modeOptions).
		Text("提现名称", "paymentInfo.name", false).
		EditText("账户名称", "name", false).
		EditText("账户姓名", "realName", false).
		EditText("卡号｜地址", "number", false).
		EditText("账户代码", "code", false).
		EditText("账户备注", "remark", false).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("操作时间", "updatedAt", true))

	return ctx.SuccessJson(config)
}
