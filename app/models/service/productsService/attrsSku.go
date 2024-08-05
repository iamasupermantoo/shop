package productsService

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"gorm.io/gorm"
	"strconv"
)

// ProductAttrsSku 产品Sku服务模型
type ProductAttrsSku struct {
	tx *gorm.DB
}

// NewProductAttrsSku 创建产品Sku服务模型
func NewProductAttrsSku(tx *gorm.DB) *ProductAttrsSku {
	return &ProductAttrsSku{tx: tx}
}

// GetStoreSkuIdOptions 获取店铺SkuId
func (_Sku *ProductAttrsSku) GetStoreSkuIdOptions(adminIds []uint) []*views.InputOptions {
	data := make([]*views.InputOptions, 0)

	// 获取管理员对应的店铺ID和父级ID都不为0的产品
	productList := make([]*productsModel.Product, 0)
	_Sku.tx.Model(&productsModel.Product{}).
		Where("admin_id IN ?", adminIds).
		Where("store_id <> ?", 0).
		Where("parent_id <> ?", 0).
		Find(&productList)

	// 获取产品对应的Sku名称
	skuList := make([]*productsModel.ProductAttrsSku, 0)
	for _, product := range productList {
		_Sku.tx.Where("product_id = ?", product.ID).
			Find(&skuList)

		for _, sku := range skuList {
			data = append(data, &views.InputOptions{
				Label: "ID: { " + strconv.Itoa(int(product.ID)) + " }，productName: { " + product.Name + "}，skuName: { " + sku.Name + " } ",
				Value: sku.ID,
			})
		}
	}

	return data
}

// GenerateProductSkuList 生成产品属性
func (_Sku *ProductAttrsSku) GenerateProductSkuList(productId uint) []productsModel.ProductAttrsSkuList {
	attrs := _Sku.getProductAttrsList(productId)
	skuList := make([]productsModel.ProductAttrsSkuList, 0)

	for i := 0; i < len(attrs[0].Values); i++ {
		sku := _Sku.generateDescartesSkuList(
			attrs,
			productsModel.ProductAttrsSkuList{
				Ids:  []string{strconv.FormatInt(int64(attrs[0].Values[i].ID), 10)},
				Name: []string{attrs[0].Values[i].Name},
			},
			1)
		skuList = append(skuList, sku...)
	}
	return skuList
}

// 获取产品属性健和值
func (_Sku *ProductAttrsSku) getProductAttrsList(productId uint) []*productsModel.ProductAttrsKeyVal {
	attrsList := make([]*productsModel.ProductAttrsKeyVal, 0)
	if _Sku.tx == nil {
		_Sku.tx = database.Db
	}

	_Sku.tx.Model(&productsModel.ProductAttrsKey{}).Where("product_id = ?", productId).
		Preload("Values").
		Find(&attrsList)
	return attrsList
}

// generateDescartesSkuList 产品属性笛卡尔积生成
func (_Sku *ProductAttrsSku) generateDescartesSkuList(attrsKeyList []*productsModel.ProductAttrsKeyVal, sep productsModel.ProductAttrsSkuList, index int) []productsModel.ProductAttrsSkuList {
	skuList := make([]productsModel.ProductAttrsSkuList, 0)

	//	如果只有一位属性,  直接返回
	if len(attrsKeyList) < 2 {
		return []productsModel.ProductAttrsSkuList{sep}
	}

	sepTmp := sep
	for i := 0; i < len(attrsKeyList[index].Values); i++ {
		sep = sepTmp
		sep.Ids = append(sep.Ids, strconv.FormatInt(int64(attrsKeyList[index].Values[i].ID), 10))
		sep.Name = append(sep.Name, attrsKeyList[index].Values[i].Name)

		if len(attrsKeyList)-1 == index {
			skuList = append(skuList, sep)
		} else {
			skuList = append(skuList, _Sku.generateDescartesSkuList(attrsKeyList, sep, index+1)...)
		}
	}

	return skuList
}
