package userBill

import (
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/wallets/bill"
)

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	assetsOptions := walletsService.NewWalletsAssets().AdminAssetsOptions(ctx.GetAdminChildIds())
	walletBillOptions := walletsService.NewUserBill().GetViewsOptions(ctx.Rds, ctx.AdminSettingId)

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Select("资产名称", "assetsId", assetsOptions).
		Select("账单类型", "type", walletBillOptions).
		RangeDatePicker("账单时间", "createdAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "id", Field: "ids", Scan: "id"}),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Select("账单类型", "type", true, walletBillOptions).
		Text("账单金额", "money", true).
		Text("用户余额", "balance", true).
		DatePicker("账单时间", "createdAt", true))

	return ctx.SuccessJson(config)
}
