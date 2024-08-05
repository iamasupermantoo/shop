package userOrder

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// SubmitIndexParams 提交页面参数
type SubmitIndexParams struct {
	SkuList []*AttrInfo `json:"skuList"`
}

type Data struct {
	StoreInfo *shopsModel.Store `json:"storeInfo"`
	OrderInfo []*OrderInfo      `json:"orderInfo"`
}

type OrderInfo struct {
	productsModel.ProductAttrsSku `gorm:"foreignKey:SkuId"`
	ProductInfo                   productsModel.Product `json:"productInfo" gorm:"foreignKey:ProductId"`
	Nums                          int                   `json:"nums" gorm:"-"`   // 购买数量
	CartId                        uint                  `json:"cartId" gorm:"-"` // 购物车Id
}

func (*OrderInfo) TableName() string {
	return "product_attrs_sku"
}

// SubmitIndex 提交订单页面信息
func SubmitIndex(ctx *context.CustomCtx, params *SubmitIndexParams) error {
	skuDataMap := make(map[uint]*AttrInfo)
	skuIds := make([]uint, 0)
	for _, v := range params.SkuList {
		skuIds = append(skuIds, v.SkuId)
		skuDataMap[v.SkuId] = v
	}

	productAndSkuList := make([]*OrderInfo, 0)
	database.Db.Preload("ProductInfo", database.Db.Where("status = ?", productsModel.ProductStatusActive)).
		Where("id IN ?", skuIds).
		Where("status = ?", productsModel.ProductAttrsSkuStatusActivate).
		Find(&productAndSkuList)
	if len(productAndSkuList) == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	storeIs := make([]uint, 0)
	storeOrderMap := make(map[uint][]*OrderInfo)
	for _, v := range productAndSkuList {
		storeId := v.ProductInfo.StoreId
		if _, ok := storeOrderMap[storeId]; !ok {
			storeIs = append(storeIs, storeId)
		}
		if p, ok := skuDataMap[v.ID]; ok {
			v.Nums = p.Nums
			v.CartId = p.CartId
		}

		storeOrderMap[storeId] = append(storeOrderMap[storeId], v)
	}

	storeList := make([]*shopsModel.Store, 0)
	database.Db.Where("id IN ?", storeIs).Where("status > ?", shopsModel.StoreStatusDisabled).Find(&storeList)

	data := make([]Data, 0)
	for _, store := range storeList {
		data = append(data, Data{
			StoreInfo: store,
			OrderInfo: storeOrderMap[store.ID],
		})
	}

	return ctx.SuccessJson(data)
}
