package productsModel

import (
	"gofiber/app/models/model/types"
)

// ProductBrowsing 店铺浏览记录
type ProductBrowsing struct {
	types.GormModel
	AdminId   uint   `json:"adminId" gorm:"type:int unsigned not null;default:0;comment:管理员ID"`
	UserId    uint   `json:"userId" gorm:"type:int unsigned not null;default:0;comment:用户ID"`
	StoreId   uint   `json:"storeId" gorm:"type:int unsigned not null;default:0;comment:店铺ID"`
	ProductId uint   `json:"productId" gorm:"type:int unsigned not null;default:0;comment:商品ID"`
	Type      int    `json:"type" gorm:"type:tinyint not null;default:1;comment:类型1默认类型"`
	Status    int    `json:"status" gorm:"type:tinyint not null;default:10;comment:状态10默认状态"`
	Data      string `json:"data" gorm:"type:varchar(255) not null;default:'';comment:数据"`
}
