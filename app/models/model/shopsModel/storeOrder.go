package shopsModel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"gofiber/app/models/model/types"
)

const (
	// DataIndexTypeUser 数据用户列表
	DataIndexTypeUser = 1
	// DataIndexTypeStore 数据店铺列表
	DataIndexTypeStore = 2

	// ProductStoreOrderStatusDisable 订单取消
	ProductStoreOrderStatusDisable = -1
	// ProductStoreOrderStatusPending 待支付
	ProductStoreOrderStatusPending = 10
	// ProductStoreOrderStatusShipping 待发货
	ProductStoreOrderStatusShipping = 12
	// ProductStoreOrderStatusProgress 待收货
	ProductStoreOrderStatusProgress = 14
	// ProductStoreOrderStatusComplete 订单完成
	ProductStoreOrderStatusComplete = 20
)

// ProductStoreOrder  店铺订单包裹
type ProductStoreOrder struct {
	types.GormModel
	AdminId     uint         `json:"adminId" gorm:"type:int unsigned not null;common:管理员ID"`
	AssetsId    uint         `json:"assetsId" gorm:"type:int unsigned not null;common:资产ID"`
	StoreId     uint         `json:"storeId" gorm:"type:int unsigned not null;default:0;comment:店铺ID"`
	UserId      uint         `json:"userId" gorm:"type:int unsigned not null;common:用户ID"`
	OrderSn     string       `json:"orderSn" gorm:"type:varchar(64);not null;index;common:订单编号"`
	Type        int          `json:"type" gorm:"type:int unsigned not null;default:1;common:类型 1商家订单"`
	Money       float64      `json:"money" gorm:"type:decimal(20,6) not null;common:购买总价"`
	FinalMoney  float64      `json:"finalMoney" gorm:"type:decimal(10,2) not null;default:0;comment:实际总价"`
	Earnings    float64      `json:"earnings" gorm:"type:decimal(10,2) not null;default:0;comment:收益总额"`
	Status      int          `json:"status" gorm:"type:tinyint not null;default:10;comment:状态 -1取消 10待支付 12待发货 14待收货 20完成"`
	AddressData *AddressData `json:"addressData" gorm:"type:varchar(2040) not null;default:'';comment:地址信息"`
	Data        string       `json:"data" gorm:"type:text;common:数据"`
}

type AddressData ShippingAddress

// Scan 查询数据
func (_AddressData *AddressData) Scan(val any) error {
	bates, ok := val.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan SkuData value:", val))
	}
	if len(bates) > 0 {
		return json.Unmarshal(bates, _AddressData)
	}
	*_AddressData = AddressData{}
	return nil
}

// Value 设置数据
func (_AddressData *AddressData) Value() (driver.Value, error) {
	if _AddressData == nil {
		addressData := AddressData{}
		return json.Marshal(&addressData)
	}
	return json.Marshal(_AddressData)
}
