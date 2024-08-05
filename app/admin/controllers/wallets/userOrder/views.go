package userOrder

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
	"strconv"
)

const (
	baseURL = "/auth/wallets/order"
)

var statusOptions = []*views.InputOptions{
	{Label: "拒绝", Value: walletsModel.WalletUserOrderStatusRefuse},
	{Label: "同意", Value: walletsModel.WalletUserOrderStatusComplete},
	{Label: "审核", Value: walletsModel.WalletUserOrderStatusActive},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	indexURL := baseURL + "/index"
	walletsUserOrderType := ctx.QueryInt("type")

	switch walletsUserOrderType {
	case walletsModel.WalletUserOrderTypeDeposit:
		indexURL = baseURL + "/deposit/balance"
	case walletsModel.WalletUserOrderTypeAssetsDeposit:
		indexURL = baseURL + "/deposit/assets"
	case walletsModel.WalletUserOrderTypeWithdraw:
		indexURL = baseURL + "/withdraw/balance"
	case walletsModel.WalletUserOrderTypeAssetsWithdraw:
		indexURL = baseURL + "/withdraw/assets"
	}

	// 创建视图
	config := views.NewTableViews(indexURL, baseURL+"/update")

	// 查询设置
	searchInputViews := views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("订单编号", "orderSn")

	// 如果资产类型添加筛选资产
	if walletsUserOrderType == walletsModel.WalletUserOrderTypeAssetsDeposit || walletsUserOrderType == walletsModel.WalletUserOrderTypeAssetsWithdraw {
		assetsOptions := walletsService.NewWalletsAssets().AdminAssetsOptions(ctx.GetAdminChildIds())
		searchInputViews.Select("资产列表", "assetsId", assetsOptions)
	}
	config.SetSearch(searchInputViews)

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "orderSn", Field: "ids", Scan: "id"}),
		},
	}...)

	// 数据表格
	columnsViews := views.NewColumnsViews().
		Text("ID", "id", true).
		Text("订单编号", "orderSn", false).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false)

	// 充值显示充值的账户名称
	if walletsUserOrderType == walletsModel.WalletUserOrderTypeDeposit || walletsUserOrderType == walletsModel.WalletUserOrderTypeAssetsDeposit {
		columnsViews.Text("充值名称", "paymentInfo.name", false)
	}

	if walletsUserOrderType == walletsModel.WalletUserOrderTypeWithdraw || walletsUserOrderType == walletsModel.WalletUserOrderTypeAssetsWithdraw {
		columnsViews.Text("提现名称", "accountInfo.name", false)
		columnsViews.Text("提现卡号｜地址", "accountInfo.number", false)
	}

	if walletsUserOrderType == walletsModel.WalletUserOrderTypeAssetsDeposit || walletsUserOrderType == walletsModel.WalletUserOrderTypeAssetsWithdraw {
		columnsViews.Text("资产名称", "assetsInfo.name", false)
	}

	if walletsUserOrderType == walletsModel.WalletUserOrderTypeDeposit || walletsUserOrderType == walletsModel.WalletUserOrderTypeWithdraw {
		columnsViews.Text("汇率金额", "rateMoney", false)
	}
	columnsViews.Text("订单金额", "money", true).
		Text("手续费", "fee", true)
	if walletsUserOrderType == walletsModel.WalletUserOrderTypeDeposit || walletsUserOrderType == walletsModel.WalletUserOrderTypeAssetsDeposit {
		columnsViews.Image("充值凭证", "voucher", false)
	}
	columnsViews.Select("状态", "status", true, statusOptions).
		DatePicker("申请时间", "createdAt", true).
		DatePicker("完成时间", "updatedAt", true)
	config.SetColumn(columnsViews)

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "同意", Color: views.ColorPositive, Size: "xs", Eval: "scope.row.status == " + strconv.Itoa(walletsModel.WalletUserOrderStatusActive)},
			Config:      views.NewDialogViews("agree", baseURL+"/agree", "同意当前请求操作").SetSizingSmall(),
		},
		{
			ButtonViews: views.ButtonViews{Label: "拒绝", Color: views.ColorNegative, Size: "xs", Eval: "scope.row.status == " + strconv.Itoa(walletsModel.WalletUserOrderStatusActive)},
			Config: views.NewDialogViews("agree", baseURL+"/refuse", "拒绝当前请求操作").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Text("拒绝理由", "data"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
