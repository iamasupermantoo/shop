package userOrder

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/shopsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type CreateParams struct {
	AssetsId  uint        `json:"assetsId"`                                  // 资产ID
	AddressId uint        `json:"addressId" validate:"required,gt=0"`        // 收货地址Id
	SkuList   []*AttrInfo `json:"skuList" validate:"required,dive,required"` // sku 信息
}

type ProductAndSkuInfo struct {
	productsModel.ProductAttrsSku
	ProductInfo *productsModel.Product `json:"productInfo" gorm:"foreignKey:ProductId;references:ID"`
}

func (*ProductAndSkuInfo) TableName() string {
	return "product_attrs_sku"
}

type AttrInfo struct {
	CartId uint `json:"cartId"`
	SkuId  uint `json:"skuId"` // skuId
	Nums   int  `json:"nums"`  // 购买数量
}

// Create 创建订单
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	userId := ctx.UserId
	// 获取收获地址
	addressInfo := &shopsModel.ShippingAddress{}
	result := database.Db.Model(addressInfo).Where("id = ?", params.AddressId).Where("user_id = ?", userId).Where("status = ?", shopsModel.ShippingAddressStatusActivate).Find(addressInfo)
	if result.Error != nil || addressInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	var userStoreId uint
	database.Db.Model(&shopsModel.Store{}).Select("id").Where("user_id = ?", userId).Scan(&userStoreId)

	userInfo := usersModel.User{}
	result = database.Db.Where("id = ?", userId).Find(&userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	skuIds := make([]uint, 0)
	cartIds := make([]uint, 0)
	attrMap := make(map[uint]*AttrInfo)
	for _, v := range params.SkuList {
		attrMap[v.SkuId] = v
		skuIds = append(skuIds, v.SkuId)
		if v.CartId > 0 {
			cartIds = append(cartIds, v.CartId)
		}
	}

	productAndSkuList := make([]*ProductAndSkuInfo, 0)
	database.Db.Preload("ProductInfo", database.Db.Where("status = ?", productsModel.ProductStatusActive)).
		Where("id IN ?", skuIds).
		Where("status = ?", productsModel.ProductAttrsSkuStatusActivate).
		Find(&productAndSkuList)
	if len(productAndSkuList) == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	skuParams := make(map[uint][]*shopsService.SkuParam)
	for _, v := range productAndSkuList {
		storeId := v.ProductInfo.StoreId
		if storeId == userStoreId {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}
		attrInfo := attrMap[v.ID]
		skuParams[storeId] = append(skuParams[storeId], &shopsService.SkuParam{
			SkuInfo: &v.ProductAttrsSku,
			Nums:    attrInfo.Nums,
		})
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		for storeId, v := range skuParams {
			err := shopsService.NewStoreOrder(tx).CreatOrder(ctx.Rds, params.AssetsId, storeId, &userInfo, addressInfo, v)
			if err != nil {
				return err
			}
		}

		// 删除购物
		if len(cartIds) > 0 {
			if err := tx.Where("id  IN ?", cartIds).Delete(&shopsModel.StoreCart{}).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	return ctx.SuccessJsonOK()
}
