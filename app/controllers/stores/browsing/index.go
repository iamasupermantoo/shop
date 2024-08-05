package browsing

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	Pagination *scopes.Pagination `json:"pagination"` // 分页数据
}

type browsingData struct {
	productsModel.ProductBrowsing
	StoreInfo   shopsModel.Store      `json:"storeInfo" gorm:"foreignKey:StoreId;"`
	ProductInfo productsModel.Product `json:"productInfo" gorm:"foreignKey:ProductId;"`
}

func (browsingData) TableName() string {
	return "product_browsing"
}

// Index 用户浏览记录
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*browsingData, 0)}
	database.Db.Model(&productsModel.ProductBrowsing{}).Where("user_id = ?", ctx.UserId).
		Preload("StoreInfo").Preload("ProductInfo").
		Count(&data.Count).Scopes(params.Pagination.Scopes()).Find(&data.Items)

	return ctx.SuccessJson(data)
}
