package userConvert

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURl = "/auth/wallets/convert"
)

var statusOptions = []*views.InputOptions{
	{Label: "完成", Value: walletsModel.WalletUserConvertStatusActive},
	{Label: "失败", Value: walletsModel.WalletUserConvertStatusDisable},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	assetsOptions := walletsService.NewWalletsAssets().AdminAssetsOptions(ctx.GetAdminChildIds())

	// 创建视图
	config := views.NewTableViews(baseURl+"/index", baseURl+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Select("发送资产", "assetsId", assetsOptions).
		Select("接收资产", "toAssetsId", assetsOptions).
		Select("状态", "status", statusOptions))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURl+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "id", Field: "ids", Scan: "id"}),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Select("发送资产", "assetsId", true, assetsOptions).
		Select("接收资产", "toAssetsId", true, assetsOptions).
		Text("转换金额", "money", true).
		Text("获取数量", "nums", true).
		Text("手续费", "fee", true).
		Select("状态", "status", true, statusOptions).
		DatePicker("完成时间", "updatedAt", true))

	return ctx.SuccessJson(config)
}
