package store

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseurl        = "/auth/shops/store"
	indexUrl       = baseurl + "/index"
	updateUrl      = baseurl + "/update"
	deleteUrl      = baseurl + "/delete"
	createOrderUrl = baseurl + "/order/create"
)

var (
	statusOptions = []*views.InputOptions{
		{Label: "开店", Value: shopsModel.StoreStatusActivate},
		{Label: "关店", Value: shopsModel.StoreStatusDisabled},
	}
	typeOptions = []*views.InputOptions{
		{Label: "普通店铺", Value: shopsModel.StoreTypeDefault},
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
		Text("联系方式", "contact").
		Text("关键词", "keywords").
		Select("店铺状态", "status", statusOptions).
		Select("店铺类型", "type", typeOptions).
		RangeDatePicker("更新时间", "updatedAt").
		GetInputListInfo()

	// 头部操作按钮
	config.Table.Tools = []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", deleteUrl, "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "Ids", Scan: "id"}),
		},
	}

	// 数据表格
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Select("店铺类型", "type", true, typeOptions).
		EditText("店铺名称", "name", false).
		Image("店铺Logo", "logo", false).
		EditText("联系方式", "contact", false).
		EditTextArea("店铺关键词", "keywords", false).
		EditNumber("评分", "rating", true).
		EditNumber("信用分", "score", true).
		EditToggle("店铺状态", "status", true, statusOptions).
		DatePicker("更新时间", "updatedAt", true).
		GetColumnsListInfo()

	// 数据操作栏目
	config.Table.Options = []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", updateUrl, "更新数据信息").
				SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Image("店铺Logo", "logo").
						Text("店铺名称", "name").
						Text("联系方式", "contact").
						Text("店铺关键词", "keywords").
						Editor("店铺描述", "desc"),
				),
		},
		{
			ButtonViews: views.ButtonViews{Label: "自动下单", Color: views.ColorSecondary, Size: "xs"},
			Config: views.NewDialogViews("crateOrder", createOrderUrl, "自动下单").
				SetSizingSmall().
				SetInputViews(
					views.NewInputViews().
						Number("购买多少个产品", "productNumber").
						Number("每个产品的购买数量", "buyNumber"),
				),
		},
	}

	return ctx.SuccessJson(config)
}
