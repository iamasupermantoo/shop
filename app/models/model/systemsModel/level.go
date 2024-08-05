package systemsModel

import "gofiber/app/models/model/types"

const (
	// LevelStatusActive 开启
	LevelStatusActive = 10
	// LevelStatusDisable 禁用
	LevelStatusDisable = -1

	// LevelTypeFullPrice 全额购买
	LevelTypeFullPrice = 1
	// LevelTypeDifference 差价购买
	LevelTypeDifference = 2
)

// Level 系统等级配置
type Level struct {
	types.GormModel
	AdminId  uint    `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	Name     string  `gorm:"type:varchar(60) not null;comment:名称" json:"name"`
	Icon     string  `gorm:"type:varchar(60) not null;comment:图标" json:"icon"`
	Symbol   int     `gorm:"type:tinyint not null;comment:标识" json:"symbol"`
	Money    float64 `gorm:"type:decimal(12,2) not null;comment:金额" json:"money"`
	Days     int     `gorm:"type:tinyint not null;comment:天数" json:"days"`
	Increase float64 `gorm:"type:decimal(12,2) not null;default:10;comment:涨幅" json:"increase"`
	Type     int     `gorm:"type:smallint not null;default:1;comment:1全额购买 2差价购买" json:"type"`
	Status   int     `gorm:"type:smallint not null;default:10;comment:状态 -1禁用 10开启" json:"status"`
	Data     string  `gorm:"type:text;comment:数据" json:"data"`
	Desc     string  `gorm:"type:text;comment:详情" json:"desc"`
}

// SystemLevelInfo 管理系统等级信息
type SystemLevelInfo struct {
	ID     uint    `json:"id"`     // ID
	Icon   string  `json:"icon"`   // 图标
	Name   string  `json:"name"`   // 名称
	Symbol int     `json:"symbol"` // 标识
	Money  float64 `json:"money"`  // 金额
	Days   int     `json:"days"`   // 天数
	Desc   string  `json:"desc"`   // 详情
}
