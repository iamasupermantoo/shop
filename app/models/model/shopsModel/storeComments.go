package shopsModel

import "gofiber/app/models/model/types"

const (
	StoreCommentsStatusPending  = 10 // 待评论
	StoreCommentsStatusComplete = 20 //	已评论
)

// StoreComment 店铺商品评论
type StoreComment struct {
	types.GormModel
	AdminId   uint              `json:"adminId" gorm:"type:int unsigned not null;default:0;comment:管理员ID"`
	UserId    uint              `json:"userId" gorm:"type:int unsigned not null;default:0;comment:用户ID"`
	StoreId   uint              `json:"storeId" gorm:"type:int unsigned not null;default:0;comment:店铺ID"`
	ProductId uint              `json:"productId" gorm:"type:int unsigned not null;default:0;comment:商品ID"`
	OrderId   uint              `json:"orderId" gorm:"type:int unsigned not null;default:0;comment:订单ID"`
	Name      string            `json:"name" gorm:"type:varchar(512) not null;comment:评论内容"`
	Rating    float64           `json:"rating" gorm:"type:decimal(3,2) not null;default:0;comment:评分"`
	Status    int               `json:"status" gorm:"type:tinyint not null;default:10;comment:状态 10待评论 20已评论"`
	Images    types.GormStrings `json:"images" gorm:"type:varchar(2048) not null;default:'';comment:买家秀"`
	Data      string            `json:"data" gorm:"type:varchar(2048) not null;default:'';comment:数据"`
}
