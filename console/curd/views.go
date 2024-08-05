package curd

var viewsTmp = `package {package}

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/app/module/views"
	"gofiber/utils"
)

// Views 视图配置
func Views(ctx *fiber.Ctx) error {
	indexURL := "/auth/{modelsPackage}/{package}/index"
	createURL := "/auth/{modelsPackage}/{package}/create"
	updateURL := "/auth/{modelsPackage}/{package}/update"
	deleteURL := "/auth/{modelsPackage}/{package}/delete"

	// 创建视图
	config := views.NewTableViews(indexURL, updateURL)

	// 查询设置
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		GetInputListInfo()

	// 头部操作按钮
	config.Table.Tools = []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", deleteURL, "批量删除选择的数据").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "ID", Field: "Ids", Scan: "ID"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增数据", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", createURL, "新增数据").
				SetSizingSmall().SetInputViews(
				views.NewInputViews(),
			),
		},
	}

	// 数据表格
	config.Table.Columns = views.NewColumnsViews().
		GetColumnsListInfo()

	// 数据操作栏目
	config.Table.Options = []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", updateURL, "更新数据信息").SetSizingSmall().
				SetInputViews(
					views.NewInputViews(),
				),
		},
	}

	return ctx.JSON(utils.SuccessJson(config))
}`
