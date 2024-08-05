package userAssets

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/users/assets"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: walletsModel.WalletUserAssetsStatusActive},
	{Label: "禁用", Value: walletsModel.WalletUserAssetsStatusDisable},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	assetsOptions := walletsService.NewWalletsAssets().AdminAssetsOptions(ctx.GetAdminChildIds())

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Select("资产名称", "assetsId", assetsOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("操作时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "用户资产加减款", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("money", baseURL+"/money", "用户资产加减款").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					SelectDefault("资产名称", "assetsId", assetsOptions).
					Text("用户名称", "userName").
					Number("操作金额(含正负)", "money").SetValue("money", 100),
			),
		},
	}...)
	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Text("资产名称", "assetsInfo.name", false).
		Text("金额", "money", true).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("操作时间", "updatedAt", true))
	return ctx.SuccessJson(config)
}
