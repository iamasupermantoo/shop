package user

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/usersService"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseURL = "/auth/users/user"
)

var userTypeOptions = []*views.InputOptions{
	{Label: "普通用户", Value: usersModel.UserTypeDefault},
	{Label: "虚拟用户", Value: usersModel.UserTypeVirtual},
	{Label: "会员用户", Value: usersModel.UserTypeLevel},
}

var statusOptions = []*views.InputOptions{
	{Label: "激活", Value: usersModel.UserStatusActive},
	{Label: "禁用", Value: usersModel.UserStatusDisable},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	childrenManageOptions := adminsService.NewAdminUser(ctx.Rds, ctx.AdminSettingId).GetAdminChildrenOptions(ctx.GetAdminChildIds())

	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("上级账户", "parentName").
		Text("账户", "userName").
		Text("昵称", "nickName").
		Text("邮箱", "email").
		Text("手机号码", "telephone").
		Select("类型", "type", userTypeOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("注册时间", "createdAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "userName", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "余额加减款", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("money", baseURL+"/money", "用户余额加减款").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().Text("用户名称", "userName").
					Number("操作金额(含正负)", "money").SetValue("money", 100),
			),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增用户", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增用户数据").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					SelectDefault("用户类型", "type", userTypeOptions).
					Text("账户", "userName").
					Password("密码", "password"),
			),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增虚拟用户", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("createVirtual", baseURL+"/virtual/create", "新增虚拟用户").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Number("虚拟用户个数", "number").
					Password("统一密码", "password"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("上级用户", "parentInfo.userName", false).
		//Image("头像", "avatar", false).
		Select("类型", "type", true, userTypeOptions).
		Text("账户", "userName", false).
		EditNumber("信用分", "score", true).
		Text("余额", "money", true).
		//EditText("昵称", "nickName", false).
		//EditText("邮箱", "email", false).
		//EditText("手机号码", "telephone", false).
		EditToggle("状态", "status", true, statusOptions).
		DatePicker("活跃时间", "updatedAt", true).
		DatePicker("注册时间", "createdAt", true))

	// 获取用户默认配置
	userSettingInputParams, userSettingInputViews := usersService.NewUserSetting(ctx.Rds, 0, 0).GetDefaultInputViews()
	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新用户信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Image("头像", "avatar").
						Select("所属管理", "adminId", childrenManageOptions).
						Number("上级用户ID", "parentId").
						Select("类型", "type", userTypeOptions).
						Password("登录密码", "password").
						Password("安全密钥", "securityKey").
						TextArea("个性签名", "desc"),
				),
		},
		{
			ButtonViews: views.ButtonViews{Label: "配置", Color: views.ColorSecondary, Size: "xs"},
			Config: views.NewDialogViews("update", "/auth/users/setting/update", "用户配置").
				SetInputViews(
					views.NewInputViews().
						Json("设置", "settingJson", userSettingInputViews.GetInputListColumn()).
						SetReadonly("settingJson").SetValue("settingJson", userSettingInputParams),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
