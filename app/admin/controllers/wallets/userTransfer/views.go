package userTransfer

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/wallets/transfer"
)

var typeOptions = []*views.InputOptions{
	{Label: "余额类型", Value: walletsModel.WalletUserTransferTypeBalance},
	{Label: "资产类型", Value: walletsModel.WalletUserTransferTypeAssets},
}
var statusOptions = []*views.InputOptions{
	{Label: "完成", Value: walletsModel.WalletUserTransferStatusActive},
	{Label: "失败", Value: walletsModel.WalletUserTransferStatusDisable},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	assetsOptions := walletsService.NewWalletsAssets().AdminAssetsOptions(ctx.GetAdminChildIds())

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("发送账户", "senderName").
		Text("接收账户", "receiverName").
		Select("资产类型", "assetsId", assetsOptions).
		Select("类型", "type", typeOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("完成时间", "updatedAt"))

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
		Text("发送账户", "senderInfo.userName", false).
		Text("接收账户", "receiverInfo.userName", false).
		Select("类型", "type", true, typeOptions).
		Select("资产类型", "assetsId", true, assetsOptions).
		Text("转移金额", "money", true).
		Text("手续费", "fee", true).
		Select("状态", "status", true, statusOptions).
		DatePicker("完成时间", "createdAt", true))

	return ctx.SuccessJson(config)
}
