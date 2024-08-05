package settled

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
	"strconv"
)

const (
	baseUrl   = "/auth/shops/settled"
	indexUrl  = baseUrl + "/index"
	createUrl = baseUrl + "/create"
	updateUrl = baseUrl + "/update"
	deleteUrl = baseUrl + "/delete"
	statusUrl = baseUrl + "/status"
)

var (
	statusOptions = []*views.InputOptions{
		{Label: "拒绝", Value: shopsModel.StoreSettledStatusRefuse},
		{Label: "审核", Value: shopsModel.StoreSettledStatusPending},
		{Label: "通过", Value: shopsModel.StoreSettledStatusPass},
	}
	typeOptions = []*views.InputOptions{
		{Label: "身份证", Value: shopsModel.StoreSettledTypeIdCard},
		{Label: "营业执照", Value: shopsModel.StoreSettledTypeLicense},
	}
)

// Views 视图配置
func Views(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	// 创建视图
	config := views.NewTableViews(indexUrl, updateUrl)

	// 查询设置
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("店铺名称", "name").
		Text("证件号码", "number").
		Text("联系方式", "contact").
		Text("邮箱地址", "email").
		Select("状态", "status", statusOptions).
		Select("类型", "type", typeOptions).
		RangeDatePicker("更新时间", "updatedAt").
		GetInputListInfo()

	// 头部操作按钮
	config.Table.Tools = []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", deleteUrl, "批量删除店铺认证数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "id", Field: "Ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "店铺申请", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", createUrl, "新增店铺申请").
				SetSizingSmall().SetInputViews(
				views.NewInputViews().
					Image("证件照正面", "photo1").
					Image("证件照背面", "photo2").
					Text("店铺名称", "name").
					Text("用户账户", "userName").
					Select("证件类型", "type", typeOptions).
					SetValue("type", shopsModel.StoreSettledTypeLicense).
					Text("证件号码", "number").
					Text("邮箱地址", "email").
					Text("联系方式", "contact"),
			),
		},
	}

	// 数据表格
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Select("证件类型", "type", true, typeOptions).
		EditText("店铺名称", "name", false).
		Image("证件照正面", "photo1", false).
		Image("证件照反面", "photo2", false).
		Image("手持证件照", "photo3", false).
		EditText("证件号码", "number", false).
		EditText("联系方式", "contact", false).
		EditText("邮箱地址", "email", false).
		EditTextArea("拒绝理由", "data", false).
		Select("入驻状态", "status", true, statusOptions).
		DatePicker("更新时间", "updatedAt", true).
		GetColumnsListInfo()

	// 数据操作栏目
	config.Table.Options = []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", updateUrl, "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Image("证件照正面", "photo1").
						Image("证件照背面", "photo2").
						Image("手持证件照", "photo3").
						Select("证件类型", "type", typeOptions).
						Text("店铺名称", "name").
						Text("证件号码", "number").
						Text("联系方式", "contact").
						Text("邮箱地址", "email").
						TextArea("数据", "data"),
				),
		},
		{
			ButtonViews: views.ButtonViews{Label: "通过", Color: views.ColorPositive, Size: "xs", Eval: "scope.row.status == " + strconv.Itoa(shopsModel.StoreSettledStatusPending)},
			Config: views.NewDialogViews("update", statusUrl, "店铺审核通过").SetSizingSmall().
				SetParams(map[string]interface{}{"opStatus": shopsModel.StoreSettledStatusPass}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "拒绝", Color: views.ColorNegative, Size: "xs", Eval: "scope.row.status == " + strconv.Itoa(shopsModel.StoreSettledStatusPending)},
			Config: views.NewDialogViews("update", statusUrl, "店铺审核拒绝").SetSizingSmall().
				SetInputViews(
					views.NewInputViews().TextArea("理由", "data"),
				),
		},
	}

	return ctx.SuccessJson(config)
}
