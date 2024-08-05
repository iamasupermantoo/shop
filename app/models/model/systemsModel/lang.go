package systemsModel

import "gofiber/app/models/model/types"

const (
	// LangStatusActive 开启
	LangStatusActive = 10

	// LangStatusDisable 禁用
	LangStatusDisable = -1
)

// Lang 系统语言
type Lang struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	Name    string `gorm:"type:varchar(60) not null;comment:名称" json:"name"`
	Alias   string `gorm:"type:varchar(60) not null;comment:别名" json:"alias"`
	Symbol  string `gorm:"type:varchar(60) not null;comment:标识" json:"symbol"`
	Icon    string `gorm:"type:varchar(60) not null;comment:图标" json:"icon"`
	Sort    int    `gorm:"type:tinyint not null;default:99;comment:排序" json:"sort"`
	Status  int    `gorm:"type:smallint not null;default:10;comment:状态 -1禁用 10开启" json:"status"`
	Data    string `gorm:"type:text;comment:数据" json:"data"`
}

// SystemLangInfo 管理系统语言信息
type SystemLangInfo struct {
	Id     uint   `json:"id"`                       //	ID
	Icon   string `json:"icon"`                     // 	图标
	Name   string `gorm:"column:alias" json:"name"` //	别名
	Symbol string `json:"symbol"`                   // 标识
}
