package product

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Category 商品分类
func Category(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	categoryList := make([]*productsModel.Category, 0)
	database.Db.Model(&productsModel.Category{}).Where("admin_id = ?", ctx.AdminSettingId).Where("status = ?", productsModel.CategoryStatusActive).Order("sort ASC").Find(&categoryList)
	data := recursionFindCategory(categoryList, 0)

	return ctx.SuccessJson(data)
}

func recursionFindCategory(categoryList []*productsModel.Category, parentId uint) []*categoryChildren {
	data := make([]*categoryChildren, 0)

	for _, category := range categoryList {
		if category.ParentId == parentId {
			data = append(data, &categoryChildren{
				Category: category,
				Children: recursionFindCategory(categoryList, category.ID),
			})
		}
	}
	return data
}

type categoryChildren struct {
	*productsModel.Category
	Children []*categoryChildren `json:"children"`
}
