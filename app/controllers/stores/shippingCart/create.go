package shippingCart

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type CreateParams struct {
	SkuId int `json:"skuId" validate:"required"` // 产品ID
	Nums  int `json:"nums" validate:"required"`  // 购买数量
}

// Create 添加购物车
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	//	查询sku是否存在
	skuInfo := &productsModel.ProductAttrsSku{}
	if result := database.Db.
		Where("id = ?", params.SkuId).Where("status = ?", productsModel.ProductAttrsSkuStatusActivate).
		Find(skuInfo); result.RowsAffected == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	productInfo := productsModel.Product{}
	if result := database.Db.Where("id = ?", skuInfo.ProductId).
		Where("status = ?", productsModel.ProductStatusActive).
		Where("type = ?", productsModel.ProductTypeDefault).
		Where("admin_id = ?", ctx.AdminSettingId).
		Find(&productInfo); result.RowsAffected == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	storeInfo := &shopsModel.Store{}
	if result := database.Db.Where("id = ?", productInfo.StoreId).
		Where("status = ?", shopsModel.StoreStatusActivate).
		Find(storeInfo); result.RowsAffected == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	// 判断该商品是否存在，如果存在则修改购买数量，否则创建
	storeCartInfo := &shopsModel.StoreCart{}
	if result := database.Db.Where("product_id = ?", productInfo.ID).
		Where("user_id = ?", ctx.UserId).
		Where("store_id = ?", storeInfo.ID).Where("sku_id = ?", skuInfo.ID).
		Find(storeCartInfo); result.RowsAffected == 0 {
		database.Db.Create(&shopsModel.StoreCart{
			AdminId:   ctx.AdminId,
			UserId:    ctx.UserId,
			StoreId:   storeInfo.ID,
			ProductId: productInfo.ID,
			SkuId:     skuInfo.ID,
			Nums:      params.Nums,
		})
	} else {
		storeCartInfo.Nums += params.Nums
		database.Db.Updates(&storeCartInfo)
	}

	return ctx.SuccessJson(storeCartInfo)
}
