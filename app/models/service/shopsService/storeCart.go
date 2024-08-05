package shopsService

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type StoreCar struct {
}

func NewStoreCar() *StoreCar {
	return &StoreCar{}
}

// GetUserStoreCartList 获取用户购物车店铺分类
func (_StoreCar *StoreCar) GetUserStoreCartList(userId uint, cartIds []uint) []*shopsModel.CartStoreInfo {
	storeIds := make([]uint, 0)
	model := database.Db.Model(&shopsModel.StoreCart{}).Select("DISTINCT(store_id)").Where("user_id = ?", userId)
	if len(cartIds) > 0 {
		model.Where("id IN ?", cartIds)
	}
	model.Find(&storeIds)

	storeList := make([]*shopsModel.CartStoreInfo, 0)
	database.Db.Model(&shopsModel.Store{}).
		Where("id IN ?", storeIds).
		Preload("CartList", func(db *gorm.DB) *gorm.DB {
			db.Where("user_id = ?", userId)
			if len(cartIds) > 0 {
				db.Where("id IN ?", cartIds)
			}
			return db
		}).
		Preload("CartList.ProductInfo").
		Preload("CartList.SkuInfo").
		Find(&storeList)

	return storeList
}
