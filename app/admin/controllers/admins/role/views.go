package role

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/admins/role"
)

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	adminCache := adminsService.NewAdminAuth(ctx.Rds, ctx.AdminSettingId)
	rolesOptions := adminCache.GetRolesChildrenOptions()

	routerNameOptions := adminsService.NewAdminAuth(ctx.Rds, ctx.AdminSettingId).GetAdminRolesRouterSelectOptions([]string{})

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Select("父级角色", "parent", rolesOptions).
		Select("角色名称", "child", rolesOptions))

	// 头部操作按钮
	config.SetTools(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "child", Field: "ids", Scan: "id"}),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "新增角色", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增角色").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Select("父级角色", "parent", rolesOptions).
					SetValue("parent", adminsModel.AuthRoleSuperManage).
					Text("新增角色", "child"),
			),
		},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "新增权限", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/auth", "新增权限").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Text("权限名称", "name").
					Text("权限路由", "auth"),
			),
		},
	)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("上级名称", "parent", false).
		EditText("角色名称", "child", false))

	// 数据操作栏目
	config.SetOptions(
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "权限配置", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/setting", "更新权限数据配置").
				SetInputViews(
					views.NewInputViews().
						Checkbox("权限目录", "authList", routerNameOptions),
				),
		},
	)
	return ctx.SuccessJson(config)
}
