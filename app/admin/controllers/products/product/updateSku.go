package product

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

// UpdateSkuParams 更新产品属性sku
type UpdateSkuParams struct {
	ID      uint                             `json:"id" validate:"required"` // 产品ID
	SkuList []*productsModel.ProductAttrsSku `json:"skuList"`                // sku列表
}

// UpdateSku 更新产品Sku
func UpdateSku(ctx *context.CustomCtx, params *UpdateSkuParams) error {
	productInfo := &productsModel.Product{}
	result := database.Db.Model(productInfo).Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(productInfo)
	if result.Error != nil || productInfo.ID == 0 {
		return ctx.ErrorJson("找不到商品信息")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(params.SkuList); i++ {
			skuInfo := params.SkuList[i]
			result = tx.Model(&productsModel.ProductAttrsSku{}).
				Where("id = ?", skuInfo.ID).Where("product_id = ?", productInfo.ID).
				Updates(&productsModel.ProductAttrsSku{
					Image:    skuInfo.Image,
					Stock:    skuInfo.Stock,
					Money:    skuInfo.Money,
					Status:   skuInfo.Status,
					Discount: skuInfo.Discount,
				})
			if result.Error != nil {
				return ctx.ErrorJson("插入规格失败[" + result.Error.Error() + "]")
			}

			// 更新当前产品价格跟折扣
			if i == 0 {
				result = tx.Model(&productsModel.Product{}).Where("id = ?", skuInfo.ProductId).Updates(&productsModel.Product{
					Money: skuInfo.Money, Discount: skuInfo.Discount,
				})
				if result.Error != nil {
					return ctx.ErrorJson("更新商品失败[" + result.Error.Error() + "]")
				}
			}
		}
		return nil
	})
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}

	return ctx.SuccessJsonOK()
}
