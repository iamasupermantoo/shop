package menu

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/admins/menu"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: adminsModel.AdminMenuStatusActive},
	{Label: "禁用", Value: adminsModel.AdminMenuStatusDisable},
}

var vueViewsOptions = []*views.InputOptions{
	{Label: "表格模版", Value: adminsModel.AdminMenuViewsTable},
	{Label: "设置模版", Value: adminsModel.AdminMenuViewsSetting},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {

	menuOptions := adminsService.NewAdminMenu(ctx.Rds, ctx.AdminSettingId).GetMenuOptions()

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")
	config.Pagination = &views.Pagination{SortBy: "sort", Descending: false, Page: 1, RowsPerPage: 500}

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("菜单名称", "name").
		Text("菜单路由", "route").
		Select("状态", "status", statusOptions))

	// 头部操作按钮
	config.SetTools(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "新增菜单", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增管理菜单").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Select("上级菜单", "parentId", menuOptions).
					Text("菜单名称", "name"),
			),
		},
	)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("父级名称", "parentInfo.name", false).
		EditText("名称", "name", false).
		EditText("路由", "route", false).
		EditNumber("排序", "sort", true).
		EditToggle("状态", "status", true, statusOptions))

	// 数据操作栏目
	config.SetOptions(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "菜单配置", Color: views.ColorSecondary, Size: "xs"},
			Config: views.NewDialogViews("menuSetting", baseURL+"/setting", "菜单参数配置").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Text("图标(Quasar图标库)", "icon").SetAlias("icon", "data.icon").
						Text("模版配置路由", "confURL").SetAlias("confURL", "data.confURL").
						Select("模版名称", "tmp", vueViewsOptions).SetAlias("tmp", "data.tmp").
						SetValue("tmp", adminsModel.AdminMenuViewsTable),
				),
		},
	)

	return ctx.SuccessJson(config)
}
