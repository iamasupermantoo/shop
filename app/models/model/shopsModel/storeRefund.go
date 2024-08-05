package shopsModel

import "gofiber/app/models/model/types"

const (
	StoreRefundStatusRefuse   = -1 // 拒绝
	StoreRefundStatusPending  = 10 // 申请
	StoreRefundStatusComplete = 20 // 同意
)

// StoreRefund 店铺产品售后
type StoreRefund struct {
	types.GormModel
	AdminId   uint              `json:"adminId" gorm:"type:int unsigned not null;default:0;comment:管理员ID"`
	UserId    uint              `json:"userId" gorm:"type:int unsigned not null;default:0;comment:用户ID"`
	OrderId   uint              `json:"orderId" gorm:"type:int unsigned not null;default:0;comment:订单ID"`
	StoreId   uint              `json:"storeId" gorm:"type:int unsigned not null;default:0;comment:店铺ID"`
	ProductId uint              `json:"productId" gorm:"type:int unsigned not null;default:0;comment:产品ID"`
	Name      string            `json:"name" gorm:"type:varchar(512) not null;comment:申请理由"`
	Images    types.GormStrings `json:"images" gorm:"type:varchar(4096) not null;comment:凭证图片"`
	Money     float64           `json:"money" gorm:"type:decimal(12,2) not null;default:0;comment:金额"`
	Data      string            `json:"data" gorm:"type:varchar(2048) not null;default:'';comment:数据"`
	Status    int               `json:"status" gorm:"type:tinyint not null;default:10;comment:售后状态 -1拒绝 10申请 20同意"`
}
