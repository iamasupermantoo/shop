package manage

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
	"strconv"
)

const (
	baseURL = "/auth/admins/manage"
)

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: adminsModel.AdminUserStatusActive},
	{Label: "冻结", Value: adminsModel.AdminUserStatusDisable},
}

var adminDataOptions = []*views.InputOptions{
	{Label: "默认模版", Value: adminsModel.AdminDefaultTemplate},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	// 下级角色列表
	adminAuthCache := adminsService.NewAdminAuth(ctx.Rds, ctx.AdminSettingId)
	childrenOptions := adminAuthCache.GetRolesChildrenOptions()
	rolesOptions := adminAuthCache.GetRolesOptions()

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "userName").
		Select("管理角色", "role", childrenOptions).
		Text("昵称", "nickName").
		Text("邮箱", "email").
		Text("域名", "domains").
		Select("状态", "status", statusOptions).
		RangeDatePicker("过期时间", "expiredAt"))

	// 头部操作按钮
	config.SetTools(&views.DialogButtonViews{
		ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
		Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的管理").
			SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "userName", Field: "ids", Scan: "id"}),
	},
		&views.DialogButtonViews{
			ButtonViews: views.ButtonViews{Label: "新增管理", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增下级管理员").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					SelectDefault("角色", "role", childrenOptions).
					Text("账户", "userName").
					Password("密码", "password"),
			),
		},
	)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Image("头像", "avatar", false).
		Text("上级", "parentInfo.userName", false).
		Text("账户", "userName", false).
		EditText("昵称", "nickName", false).
		//EditText("邮箱", "email", false).
		//EditNumber("余额", "money", true).
		Select("角色", "roleInfo.name", false, rolesOptions).
		EditText("域名", "domains", false).
		//EditText("坐席链接", "seatLink", false).
		EditText("在线客服", "online", false).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("过期时间", "expiredAt", true))

	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "管理更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().Image("头像", "avatar").
						Password("登录密码", "password").Password("安全密钥", "securityKey").
						Select("角色", "role", rolesOptions).
						SetAlias("role", "roleInfo.name").
						Select("状态", "status", statusOptions).
						Text("邮箱", "email").
						Text("坐席链接", "seatLink").
						DatePicker("过期时间", "expiredAt"),
				),
		},
		{
			ButtonViews: views.ButtonViews{Label: "商户配置", Color: views.ColorSecondary, Size: "xs", Eval: "scope.row.parentId == " + strconv.Itoa(adminsModel.SuperAdminId)},
			Config: views.NewDialogViews("merchantSetting", baseURL+"/setting", "商户配置信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Text("授权Key", "key").SetAlias("key", "data.key").
						Select("显示模版", "template", adminDataOptions).SetAlias("template", "data.template").
						Number("下级数量", "nums").SetAlias("nums", "data.agentNums").
						TextArea("白名单", "whitelist").SetAlias("whitelist", "data.whitelist"),
				),
		},
		{
			ButtonViews: views.ButtonViews{Label: "重置配置", Color: views.ColorNegative, Size: "xs", Eval: "scope.row.parentId == " + strconv.Itoa(adminsModel.SuperAdminId)},
			Config:      views.NewDialogViews("merchantReset", baseURL+"/setting/reset", "商户配置重置").SetSizingSmall(),
		},
		{
			ButtonViews: views.ButtonViews{Label: "同步配置", Color: views.ColorNegative, Size: "xs", Eval: "scope.row.parentId == " + strconv.Itoa(adminsModel.SuperAdminId)},
			Config:      views.NewDialogViews("merchantSync", baseURL+"/setting/sync", "商户配置同步").SetSizingSmall(),
		},
		{
			ButtonViews: views.ButtonViews{Label: "初始化产品", Color: views.ColorNegative, Size: "xs", Eval: "scope.row.parentId == " + strconv.Itoa(adminsModel.SuperAdminId)},
			Config:      views.NewDialogViews("initProduct", baseURL+"/product/init", "初始化产品").SetSizingSmall(),
		},
	}...)

	return ctx.SuccessJson(config)
}
