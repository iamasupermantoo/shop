package menu

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/models/service/usersService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/systems/menu"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: systemsModel.MenuStatusActive},
	{Label: "禁用", Value: systemsModel.MenuStatusDisable},
}

var typeOptions = []*views.InputOptions{
	{Label: "导航菜单", Value: systemsModel.MenuTypeNavigation},
	{Label: "用户菜单", Value: systemsModel.MenuTypeSetting},
	{Label: "更多菜单", Value: systemsModel.MenuTypeMore},
	{Label: "商户菜单", Value: systemsModel.MenuTypeStore},
}

var boolOptions = []*views.InputOptions{
	{Label: "开启", Value: types.ModelBoolTrue},
	{Label: "关闭", Value: types.ModelBoolFalse},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	menuOptions := usersService.NewUserMenu().GetAdminOptions(ctx.GetAdminChildIds())

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Select("父级", "parentId", menuOptions).
		Select("类型", "type", typeOptions).
		Text("路由", "route").
		RangeDatePicker("时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/update", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增前台菜单", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增前台菜单").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Select("父级名称", "parentId", menuOptions).
					SelectDefault("类型", "type", typeOptions).
					Text("名称(翻译)", "name").
					Text("路由", "route"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Select("类型", "type", true, typeOptions).
		Select("父级", "parentId", true, menuOptions).
		Translate("名称(翻译)", "name", false).
		Image("图标", "icon", false).
		Image("激活图标", "activeIcon", false).
		EditText("路由", "route", false).
		EditNumber("排序", "sort", true).
		EditToggle("桌面显示", "isDesktop", true, boolOptions).
		EditToggle("手机显示", "isMobile", true, boolOptions).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("时间", "updatedAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新前台菜单数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Select("父级", "parentId", menuOptions).
						Select("类型", "type", typeOptions).
						Image("图标", "icon").
						Image("激活图标", "activeIcon"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
