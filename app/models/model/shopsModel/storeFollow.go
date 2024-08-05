package shopsModel

import (
	"gofiber/app/models/model/types"
	"gofiber/app/module/database"
)

const (
	StoreFollowTypeConcernStore      = 1 // 关注店铺
	StoreFollowTypeCollectionProduct = 2 // 收藏商品

	StoreFollowStatusCancels = -1 // 取消
	StoreFollowStatusConcern = 10 // 关注
)

// StoreFollow 店铺商品关注
type StoreFollow struct {
	types.GormModel
	AdminId   uint   `json:"adminId" gorm:"type:int unsigned not null;default:0;comment:管理员ID"`
	UserId    uint   `json:"userId" gorm:"type:int unsigned not null;default:0;comment:用户ID"`
	StoreId   uint   `json:"storeId" gorm:"type:int unsigned not null;default:0;comment:店铺ID"`
	ProductId uint   `json:"productId" gorm:"type:int unsigned not null;default:0;comment:商品ID"`
	Type      int    `json:"type" gorm:"type:tinyint not null;default:1;comment:1关注店铺 2收藏商品"`
	Status    int    `json:"status" gorm:"type:tinyint not null;default:10;comment:状态 -1取消 10关注"`
	Data      string `json:"data" gorm:"type:varchar(2048) not null;default:'';comment:数据"`
}

// IsFollow 是否收藏产品商店
func IsFollow(id, userId uint, followType int) bool {
	model := database.Db
	if followType == StoreFollowTypeConcernStore {
		model.Where("store_id = ?", id).Where("type = ?", StoreFollowTypeConcernStore)
	}

	if followType == StoreFollowTypeConcernStore {
		model.Where("product = ?", id).Where("type = ?", StoreFollowTypeCollectionProduct)
	}

	if result := model.Where("user_id = ?", userId).
		Where("product_id = ?", id).
		Where("status = ?", StoreFollowStatusConcern).
		Find(&StoreFollow{}); result.RowsAffected > 0 {
		return true
	}
	return false
}
