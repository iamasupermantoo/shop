package usersModel

import (
	"gofiber/app/models/model/types"
	"time"
)

const (
	// UserLevelStatusActive 开启
	UserLevelStatusActive = 10

	// UserLevelStatusDisable 禁用
	UserLevelStatusDisable = -1

	// BuyLevelModePriceDifference 差价购买
	BuyLevelModePriceDifference = "1"
)

// UserLevel 会员表
type UserLevel struct {
	types.GormModel
	AdminId   uint      `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	UserId    uint      `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	Name      string    `gorm:"type:varchar(60) not null;comment:名称" json:"name"`
	Icon      string    `gorm:"type:varchar(60) not null;comment:图标" json:"icon"`
	Symbol    int       `gorm:"type:tinyint not null;comment:标识" json:"symbol"`
	Money     float64   `gorm:"type:decimal(12,2) not null;comment:金额" json:"money"`
	Increase  float64   `gorm:"type:decimal(12,2) not null;comment:涨幅" json:"increase"`
	Status    int       `gorm:"type:tinyint not null;default:10;comment:状态 -1禁用 10开启" json:"status"`
	Data      string    `gorm:"type:text;comment:数据" json:"data"`
	ExpiredAt time.Time `gorm:"type:datetime(3);comment:过期时间" json:"expiredAt"`
}
