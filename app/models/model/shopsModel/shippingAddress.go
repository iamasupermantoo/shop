package shopsModel

import (
	"gofiber/app/models/model/types"
)

const (
	ShippingAddressStatusDisabled = -1 // 禁用
	ShippingAddressStatusActivate = 10 // 激活

	ShippingAddressIsShowYes = 2 // 默认地址
	ShippingAddressIsShowNo  = 1 // 非默认地址

	ShippingAddressTypeReceiving = 1 // 收货地址
	ShippingAddressTypeShipments = 2 // 发货地址
)

// ShippingAddress 用户购物地址
type ShippingAddress struct {
	types.GormModel
	AdminId uint   `json:"adminId" gorm:"type:int unsigned not null;default:0;comment:管理员ID"`
	UserId  uint   `json:"userId" gorm:"type:int unsigned not null;default:0;comment:用户ID"`
	Name    string `json:"name" gorm:"type:varchar(50) not null;default:'';comment:收件人名称"`
	Contact string `json:"contact" gorm:"type:varchar(50) not null;default:'';comment:联系方式"`
	City    string `json:"city" gorm:"type:varchar(255) not null;default:'';comment:国家城市"`
	Address string `json:"address" gorm:"type:varchar(255) not null;default:'';comment:详细地址"`
	Type    int    `json:"type" gorm:"type:tinyint not null;default:1;comment:类型 1收货地址 2发货地址"`
	Status  int    `json:"status" gorm:"type:tinyint not null;default:10;comment:状态 -1禁用 10激活"`
	IsShow  int    `json:"isShow" gorm:"type:tinyint not null;default:1;comment:1不默认 2默认"`
}
