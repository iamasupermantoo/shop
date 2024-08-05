package userAuth

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
	"strconv"
)

const (
	baseURL = "/auth/users/auth"
)

var statusOptions = []*views.InputOptions{
	{Label: "完成", Value: usersModel.UserAuthStatusComplete},
	{Label: "审核", Value: usersModel.UserAuthStatusActive},
	{Label: "拒绝", Value: usersModel.UserAuthStatusRefuse},
}

var authTypeOptions = []*views.InputOptions{
	{Label: "身份证", Value: usersModel.UserAuthTypeIdCard},
}

// Views 视图配置
func Views(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(baseURL+"/index", baseURL+"/update")

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("证件姓名", "realName").
		Text("证件号码", "number").
		Text("证件地址", "address").
		Select("状态", "status", statusOptions).
		RangeDatePicker("申请时间", "createdAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", baseURL+"/delete", "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "realName", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增用户认证", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", baseURL+"/create", "新增用户认证").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Image("证件照1", "photo1").
					Image("证件照2", "photo2").
					SelectDefault("证件类型", "type", authTypeOptions).
					Text("用户账户", "userName").
					Text("证件姓名", "realName").
					Text("证件号码", "number"),
			),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Select("证件类型", "type", true, authTypeOptions).
		Image("证件照1", "photo1", false).
		Image("证件照2", "photo2", false).
		Image("证件照3", "photo3", false).
		EditText("证件姓名", "realName", false).
		EditText("证件号码", "number", false).
		EditText("证件地址", "address", false).
		Select("状态", "status", true, statusOptions).
		DatePicker("申请时间", "createdAt", true))

	// 数据操作栏目
	config.SetOptions([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", baseURL+"/update", "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Image("证件照1", "photo1").Image("证件照2", "photo2").Image("证件照3", "photo3").
						Select("证件类型", "type", authTypeOptions).
						Select("状态", "status", statusOptions),
				),
		},
		{
			ButtonViews: views.ButtonViews{Label: "同意", Color: views.ColorPositive, Size: "xs", Eval: "scope.row.status == " + strconv.Itoa(usersModel.UserAuthStatusActive)},
			Config:      views.NewDialogViews("agree", baseURL+"/agree", "同意当前请求操作").SetSizingSmall(),
		},
		{
			ButtonViews: views.ButtonViews{Label: "拒绝", Color: views.ColorNegative, Size: "xs", Eval: "scope.row.status == " + strconv.Itoa(usersModel.UserAuthStatusActive)},
			Config: views.NewDialogViews("agree", baseURL+"/refuse", "拒绝当前请求操作").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Text("拒绝理由", "data"),
				),
		},
	}...)

	return ctx.SuccessJson(config)
}
