package productsModel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"gofiber/app/models/model/types"
	"time"
)

const (
	// ProductOrderTypeDefault 商家订单
	ProductOrderTypeDefault = 1
	// ProductOrderStatusDisable 订单取消
	ProductOrderStatusDisable = -1
	// ProductOrderStatusPending 待付款
	ProductOrderStatusPending = 10
	// ProductOrderStatusShipping 待发货
	ProductOrderStatusShipping = 12
	// ProductOrderStatusProgress 待收货
	ProductOrderStatusProgress = 14
	// ProductOrderStatusComplete 订单完成
	ProductOrderStatusComplete = 20
)

// ProductOrder  产品订单表
type ProductOrder struct {
	types.GormModel
	AdminId      uint      `json:"adminId" gorm:"type:int unsigned not null;common:管理员ID"`
	UserId       uint      `json:"userId" gorm:"type:int unsigned not null;common:用户ID"`
	AssetsId     uint      `json:"assetsId" gorm:"type:int unsigned not null;common:资产ID"`
	StoreOrderId uint      `json:"storeOrderId" gorm:"type:int unsigned not null;common:店铺订单ID"`
	StoreId      uint      `json:"storeId" gorm:"type:int unsigned not null;default:0;comment:店铺ID"`
	ProductId    uint      `json:"productId" gorm:"type:int unsigned not null;common:产品ID"`
	OrderSn      string    `json:"orderSn" gorm:"type:varchar(64);not null;common:订单编号"`
	Nums         int       `json:"nums" gorm:"type:int not null;common:产品数量"`
	Money        float64   `json:"money" gorm:"type:decimal(20,6) not null;common:购买金额"`
	FinalMoney   float64   `json:"finalMoney" gorm:"type:decimal(10,2) not null;default:0;comment:实际价格"`
	Earnings     float64   `json:"earnings" gorm:"type:decimal(10,2) not null;default:0;comment:订单收益"`
	Fee          float64   `json:"fee" gorm:"type:decimal(20,6) not null;common:手续费"`
	Type         int       `json:"type" gorm:"type:int unsigned not null;default:1;common:类型1默认类型"`
	Status       int       `json:"status" gorm:"type:tinyint not null;default:10;comment:状态-1取消 10等待 11运行 20完成"`
	SkuData      *SkuData  `json:"skuData" gorm:"type:varchar(2040) not null;default:'';comment:Sku信息"`
	Data         string    `json:"data" gorm:"type:text;common:数据"`
	ExpiredAt    time.Time `json:"expiredAt" gorm:"type:datetime(3);comment:过期时间"`
}

type SkuData ProductAttrsSku

// Scan 查询数据
func (_SkuData *SkuData) Scan(val any) error {
	bates, ok := val.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan SkuData value:", val))
	}
	if len(bates) > 0 {
		return json.Unmarshal(bates, _SkuData)
	}
	*_SkuData = SkuData{}
	return nil
}

// Value 设置数据
func (_SkuData *SkuData) Value() (driver.Value, error) {
	if _SkuData == nil {
		skuData := SkuData{}
		return json.Marshal(&skuData)
	}
	return json.Marshal(_SkuData)
}
