package product

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/service/productsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	CategoryId uint               `json:"categoryId"` // 产品分类Id
	StoreId    int                `json:"storeId"`    // 店铺Id
	Search     string             `json:"search"`     // 搜索商品详情内容
	Pagination *scopes.Pagination `json:"pagination"` // 分页数据
}

type productInfo struct {
	productsModel.Product
	FollowInfo shopsModel.StoreFollow `json:"followInfo" gorm:"foreignKey:ProductId;"`
}

func (_productInfo *productInfo) TableName() string {
	return "product"
}

// Index 产品列表接口
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	settingId := ctx.AdminSettingId
	categoryIds := make([]int, 0)
	data := &context.IndexData{Items: make([]*productInfo, 0)}
	model := database.Db.Model(&productsModel.Product{})
	if params.CategoryId > 0 {
		categoryList := make([]*productsModel.Category, 0)
		database.Db.Model(&productsModel.Category{}).
			Where("admin_id = ?", settingId).
			Where("status = ?", productsModel.CategoryStatusActive).
			Find(&categoryList)
		categoryIds = productsService.NewProductCategory(ctx.Rds, settingId).
			CategoryChildrenIds(categoryList, params.CategoryId)
	}

	model.Preload("FollowInfo", database.Db.Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.StoreFollowStatusConcern)).
		Where("status = ?", productsModel.ProductStatusActive).
		Where("admin_id = ?", settingId).
		//Where("type = ?", productsModel.ProductTypeDefault).
		Scopes(scopes.NewScopes().
			In("category_id", categoryIds).
			Eq("store_id", params.StoreId).
			Like("name", params.Search+"%").
			Scopes()).
		Scopes(params.Pagination.Scopes()).
		Count(&data.Count).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
