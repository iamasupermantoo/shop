package productsService

import (
	"errors"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/module/crawling"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

const (
	ProductAttrsSkuWholesale = "wholesale" //	店铺上架操作
	ProductAttrsSkuMerchant  = "merchant"  //	商家同步操作
)

// Product 产品服务层模型
type Product struct {
	tx *gorm.DB
}

func NewProduct(tx *gorm.DB) *Product {
	return &Product{tx: tx}
}

// WholesaleShelvesProduct 产品上架产品方法
func (_Product *Product) WholesaleShelvesProduct(opt string, sourceId, productId uint, increase float64) error {
	if _Product.tx == nil {
		_Product.tx = database.Db
	}
	// 插入产品属性
	attrsList := make([]*productsModel.ProductAttrsKeyVal, 0)
	result := database.Db.Preload("Values").Where("product_id = ?", sourceId).Find(&attrsList)
	if result.Error != nil {
		return result.Error
	}

	// 插入产品属性
	if len(attrsList) == 0 {
		return errors.New("abnormalOperation")
	}

	attrs := make([]*productsModel.ProductAttrsKeyVal, 0)
	for _, attrsKey := range attrsList {
		newAttrsKey := &productsModel.ProductAttrsKeyVal{ProductAttrsKey: productsModel.ProductAttrsKey{
			ProductId: productId,
			Name:      attrsKey.Name,
			Type:      attrsKey.Type,
			Data:      attrsKey.Data,
			Status:    attrsKey.Status,
		}, Values: make([]*productsModel.ProductAttrsVal, 0)}
		for _, attrsVal := range attrsKey.Values {
			newAttrsKey.Values = append(newAttrsKey.Values, &productsModel.ProductAttrsVal{
				Name:   attrsVal.Name,
				Data:   attrsVal.Data,
				Status: attrsVal.Status,
			})
		}
		attrs = append(attrs, newAttrsKey)
	}
	result = _Product.tx.Create(attrs)
	if result.Error != nil {
		return result.Error
	}

	// 插入sku 数据
	skuList := make([]*productsModel.ProductAttrsSku, 0)
	generateSkuList := NewProductAttrsSku(_Product.tx).GenerateProductSkuList(productId)
	for _, skuInfo := range generateSkuList {
		currentSkuInfo := &productsModel.ProductAttrsSku{}
		database.Db.Model(currentSkuInfo).Where("product_id = ?", sourceId).Where("name = ?", strings.Join(skuInfo.Name, ".")).Find(currentSkuInfo)
		if currentSkuInfo.ID > 0 {
			var skuParentId uint
			var skuMoney float64
			if opt == ProductAttrsSkuWholesale {
				skuParentId = currentSkuInfo.ID
				skuMoney = currentSkuInfo.Money + currentSkuInfo.Money*increase/100
			}

			skuList = append(skuList, &productsModel.ProductAttrsSku{
				ParentId:  skuParentId,
				ProductId: productId,
				Vals:      strings.Join(skuInfo.Ids, ","),
				Name:      strings.Join(skuInfo.Name, "."),
				Image:     currentSkuInfo.Image,
				Stock:     currentSkuInfo.Stock,
				Sales:     currentSkuInfo.Sales,
				Money:     skuMoney,
				Discount:  currentSkuInfo.Discount,
				Data:      currentSkuInfo.Data,
				Status:    currentSkuInfo.Status,
			})
		}
	}
	result = _Product.tx.Create(skuList)
	return result.Error
}

// InserterProductAttrs	插入产品数据
func (_Product *Product) InserterProductAttrs(opt string, oldProductId, newProductId uint) error {
	if _Product.tx == nil {
		_Product.tx = database.Db
	}
	productsAttrsSkuList := make([]*productsModel.ProductAttrsSku, 0)
	result := database.Db.Where("product_id = ?", oldProductId).Find(&productsAttrsSkuList)
	if result.RowsAffected == 0 {
		return errors.New("abnormalOperation")
	}

	// 获取产品属性
	keyAndValList := make([]*productsModel.ProductAttrsKeyVal, 0)
	result = database.Db.Preload("Values").Where("product_id = ?", oldProductId).Find(&keyAndValList)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("abnormalOperation")
	}

	// 获取对应sku vals值
	vals := make(map[uint]uint)
	for _, sku := range productsAttrsSkuList {
		ids := utils.StringToIntList(sku.Vals)
		for _, id := range ids {
			if _, ok := vals[uint(id)]; !ok {
				vals[uint(id)] = 0
			}
		}
	}

	for _, keyAdnVal := range keyAndValList {
		productAttrsKeyInfo := productsModel.ProductAttrsKey{
			ProductId: newProductId,
			Name:      keyAdnVal.Name,
			Type:      keyAdnVal.Type,
			Data:      keyAdnVal.Data,
			Status:    keyAdnVal.Status,
		}
		if err := _Product.tx.Create(&productAttrsKeyInfo).Error; err != nil {
			return err
		}
		for _, v := range keyAdnVal.Values {
			sourceId := v.ID
			v.ID = 0
			v.KeyId = productAttrsKeyInfo.ID
			// 创建
			if err := _Product.tx.Create(&v).Error; err != nil {
				return err
			}
			if _, ok := vals[sourceId]; ok {
				vals[sourceId] = v.ID
			}
		}
	}

	for _, sku := range productsAttrsSkuList {
		ids := utils.StringToIntList(sku.Vals)
		var newVals []string
		for _, id := range ids {
			if v, ok := vals[uint(id)]; ok {
				valStr := strconv.Itoa(int(v))
				newVals = append(newVals, valStr)
			}
		}
		if len(newVals) > 0 && newVals[0] == "0" {
			return errors.New("values len = 0")
		}
		sku.Vals = strings.Join(newVals, ",")
		sku.ProductId = newProductId
		if opt == ProductAttrsSkuWholesale {
			sku.ParentId = sku.ID
		}
		sku.ID = 0
		if err := _Product.tx.Create(&sku).Error; err != nil {
			return err
		}
	}
	return nil
}

// InsertCrawlingProduct 插入爬取的数据
func (_Product *Product) InsertCrawlingProduct(settingAdminId, categoryId uint, productAttr *crawling.ProductAttr) (uint, error) {
	if len(productAttr.Style) == 0 {
		return 0, errors.New("not find style")
	}

	discount := utils.FloatAccuracy((productAttr.OriginalPrice-productAttr.CurrentPrice)/productAttr.OriginalPrice, 2)

	if productAttr.CurrentPrice == 0.0 {
		productAttr.CurrentPrice = float64(utils.NewRandom().Intn(50, 100))
	}

	// 插入爬取产品
	productInfo := &productsModel.Product{
		AdminId:    settingAdminId,
		CategoryId: categoryId,
		Name:       productAttr.Title,
		Images:     productAttr.Images,
		Money:      productAttr.CurrentPrice,
		Discount:   discount,
		Type:       productsModel.ProductTypeWholesale,
		Desc:       productAttr.Describe,
	}
	if result := _Product.tx.Create(productInfo); result.Error != nil {
		return 0, result.Error
	}

	productAttrsKeyVal := make([]*productsModel.ProductAttrsKeyVal, 0)
	for key, values := range productAttr.Style {
		attrsKeyInfo := &productsModel.ProductAttrsKeyVal{
			ProductAttrsKey: productsModel.ProductAttrsKey{
				ProductId: productInfo.ID,
				Name:      key,
			},
		}
		for _, value := range values {
			attrsKeyInfo.Values = append(attrsKeyInfo.Values, &productsModel.ProductAttrsVal{
				Name: value,
			})
		}
		productAttrsKeyVal = append(productAttrsKeyVal, attrsKeyInfo)
	}

	// 插入产品的键值属性
	if err := _Product.tx.Create(&productAttrsKeyVal).Error; err != nil {
		return 0, err
	}

	skuService := NewProductAttrsSku(_Product.tx)
	attrs := skuService.GenerateProductSkuList(productInfo.ID)
	productSkuList := make([]*productsModel.ProductAttrsSku, 0)
	for _, skuInfo := range attrs {
		productSkuList = append(productSkuList, &productsModel.ProductAttrsSku{
			ProductId: productInfo.ID,
			Vals:      strings.Join(skuInfo.Ids, ","),
			Name:      strings.Join(skuInfo.Name, "."),
			Money:     productInfo.Money,
			Discount:  discount,
		})
	}

	// 插入产品sku属性
	if err := _Product.tx.Create(&productSkuList).Error; err != nil {
		return 0, err
	}

	return productInfo.ID, nil
}
