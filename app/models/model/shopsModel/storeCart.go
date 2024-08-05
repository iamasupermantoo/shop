package shopsModel

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/types"
)

// StoreCart 店铺购物车
type StoreCart struct {
	types.GormModel
	AdminId   uint   `json:"adminId" gorm:"type:int unsigned not null;default:0;comment:管理员ID"`
	UserId    uint   `json:"userId" gorm:"type:int unsigned not null;default:0;comment:用户ID"`
	StoreId   uint   `json:"storeId" gorm:"type:int unsigned not null;default:0;comment:店铺ID"`
	ProductId uint   `json:"productId" gorm:"type:int unsigned not null;default:0;comment:商品ID"`
	SkuId     uint   `json:"skuId" gorm:"type:int unsigned not null;default:0;comment:skuId"`
	Nums      int    `json:"nums" gorm:"type:smallint unsigned not null;default:1;comment:数量"`
	Data      string `json:"data" gorm:"type:varchar(255);default:'';comment:数据"`
}

// CartStoreInfo 购物车店铺信息
type CartStoreInfo struct {
	Store
	CartList []*CartInfo `json:"cartList" gorm:"foreignKey:StoreId"`
}

func (CartStoreInfo) TableName() string {
	return "store"
}

// CartInfo 购物车信息
type CartInfo struct {
	StoreCart
	ProductInfo productsModel.Product         `json:"productInfo" gorm:"foreignKey:ProductId"`
	SkuInfo     productsModel.ProductAttrsSku `json:"skuInfo" gorm:"foreignKey:SkuId"`
}

func (CartInfo) TableName() string {
	return "store_cart"
}
