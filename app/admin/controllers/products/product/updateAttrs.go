package product

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/service/productsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
	"strings"
)

// AttrsUpdateParams  更新产品属性
type AttrsUpdateParams struct {
	ID    uint                                `json:"id" validate:"required"`
	Attrs []*productsModel.ProductAttrsKeyVal `json:"attrs"` // 产品属性的键和值
}

// UpdateAttrs 更新产品属性
func UpdateAttrs(ctx *context.CustomCtx, params *AttrsUpdateParams) error {
	productInfo := productsModel.Product{}
	result := database.Db.Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(&productInfo)
	if result.RowsAffected == 0 {
		return ctx.ErrorJson("未找到对应的商品")
	}

	// 检查当前数据是否存在产品ID
	for _, attrKey := range params.Attrs {
		if attrKey.ProductId == 0 {
			attrKey.ProductId = productInfo.ID
		}
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		// 删除属性
		tx.Unscoped().Delete(&params.Attrs)

		result = tx.Create(&params.Attrs)
		if result.Error != nil {
			return ctx.ErrorJson("属性重置失败[" + result.Error.Error() + "]")
		}

		attrsList := make([]*productsModel.ProductAttrsKeyVal, 0)
		tx.Model(&productsModel.ProductAttrsKeyVal{}).Where("product_id = ?", productInfo.ID).
			Preload("Values").
			Find(&attrsList)

		// 删除之前的sku信息
		result = tx.Unscoped().Where("product_id = ?", productInfo.ID).Delete(&productsModel.ProductAttrsSku{})
		if result.Error != nil {
			return ctx.ErrorJson("属性规格重置失败[" + result.Error.Error() + "]")
		}

		// 新增新的sku信息
		generateAttrsSkuList := productsService.NewProductAttrsSku(tx).GenerateProductSkuList(productInfo.ID)
		skuList := make([]*productsModel.ProductAttrsSku, 0)
		for _, attrSku := range generateAttrsSkuList {
			skuList = append(skuList, &productsModel.ProductAttrsSku{
				ProductId: productInfo.ID, Vals: strings.Join(attrSku.Ids, ","), Name: strings.Join(attrSku.Name, "."),
				Image: productInfo.Images[0], Money: productInfo.Money, Discount: productInfo.Discount,
			})
		}
		result = tx.Create(skuList)
		if result.Error != nil {
			return ctx.ErrorJson("插入规格失败[" + result.Error.Error() + "]")
		}
		return nil
	})
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}

	return ctx.SuccessJsonOK()
}
