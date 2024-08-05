package shopsModel

import (
	"gofiber/app/models/model/types"
)

const (
	StoreStatusDisabled = -1 //	禁用
	StoreStatusActivate = 10 //	开店
	StoreTypeDefault    = 1  //	普通店铺
)

// Store 店铺
type Store struct {
	types.GormModel
	AdminId  uint    `json:"adminId" gorm:"type:int unsigned not null;default:0;comment:管理员ID"`
	UserId   uint    `json:"userId" gorm:"type:int unsigned not null;default:0;comment:用户ID"`
	Logo     string  `json:"logo" gorm:"type:varchar(255) not null;default:'';comment:Logo"`
	Name     string  `json:"name" gorm:"type:varchar(255) not null;default:'';comment:店铺名称"`
	Contact  string  `json:"contact" gorm:"type:varchar(50) not null;default:'';comment:联系方式"`
	Type     int     `json:"type" gorm:"type:tinyint not null;default:1;comment:类型 1普通店铺"`
	Keywords string  `json:"keywords" gorm:"type:varchar(255) not null;default:'';comment:关键词"`
	Address  string  `json:"address" gorm:"type:varchar(2048) not null;default:'';comment:店铺地址"`
	Desc     string  `json:"desc" gorm:"type:varchar(255) not null;default:'';comment:描述"`
	Rating   float64 `json:"rating" gorm:"type:decimal(3,2) not null;default:5;comment:评分"`
	Score    int     `json:"score" gorm:"type:tinyint not null;default:100;comment:信用分"`
	Sales    int     `json:"sales" gorm:"type:int unsigned not null;default:0;comment:销售量"`
	Status   int     `json:"status" gorm:"type:tinyint not null;default:10;comment:状态 -1禁用 10激活"`
	Data     string  `json:"data" gorm:"type:varchar(255) not null;default:'';comment:数据"`
}
