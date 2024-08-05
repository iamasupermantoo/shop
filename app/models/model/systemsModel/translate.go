package systemsModel

import "gofiber/app/models/model/types"

const (
	// TranslateTypeDefault 系统翻译
	TranslateTypeDefault = 1

	// TranslateTypeFrontend 前台翻译
	TranslateTypeFrontend = 2
)

// Translate 系统语言翻译表
type Translate struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	Lang    string `gorm:"type:varchar(60) not null;comment:语言标识" json:"lang"`
	Name    string `gorm:"type:varchar(60) not null;comment:名称" json:"name"`
	Type    int    `gorm:"type:tinyint unsigned not null;default:1;comment:类型 1默认翻译 2前台翻译" json:"type"`
	Field   string `gorm:"type:varchar(60) not null;comment:键名" json:"field"`
	Value   string `gorm:"type:text;comment:键值" json:"value"`
}

// SystemTranslateInfo 管理系统翻译信息
type SystemTranslateInfo struct {
	Field string `json:"label"` //	翻译名
	Value string `json:"value"` //	翻译值
}
