package follow

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/views"
)

const (
	baseUrl   = "/auth/shops/follow"
	indexUrl  = baseUrl + "/index"
	updateUrl = baseUrl + "/update"
	deleteUrl = baseUrl + "/delete"
)

var (
	typeOptions = []*views.InputOptions{
		{Label: "关注店铺", Value: shopsModel.StoreFollowTypeConcernStore},
		{Label: "收藏商品", Value: shopsModel.StoreFollowTypeCollectionProduct},
	}
	statusOptions = []*views.InputOptions{
		{Label: "关注", Value: shopsModel.StoreFollowStatusConcern},
		{Label: "取关", Value: shopsModel.StoreFollowStatusCancels},
	}
)

// Views 视图配置
func Views(ctx *context.CustomCtx, params *context.NoRequestBody) error {

	// 创建视图
	config := views.NewTableViews(indexUrl, updateUrl)

	// 查询设置
	config.SetSearch(views.NewInputViews().
		Text("管理账户", "adminName").
		Text("用户账户", "userName").
		Text("店铺名称", "storeName").
		Text("商品名称", "productName").
		Select("类型", "type", typeOptions).
		Select("状态", "status", statusOptions).
		RangeDatePicker("更新时间", "updatedAt"))

	// 头部操作按钮
	config.SetTools([]*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", deleteUrl, "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "id", Field: "Ids", Scan: "id"}),
		},
	}...)

	// 数据表格
	config.SetColumn(views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理账户", "adminInfo.userName", false).
		Text("用户账户", "userInfo.userName", false).
		Select("关注类型", "type", false, typeOptions).
		Text("店铺名称", "storeInfo.name", false).
		Text("商品名称", "productInfo.name", false).
		EditToggle("状态", "status", false, statusOptions).
		DatePicker("更新时间", "updatedAt", true))

	// 数据操作栏目
	config.Table.Options = []*views.DialogButtonViews{}

	return ctx.SuccessJson(config)
}
