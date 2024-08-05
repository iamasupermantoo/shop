package systemsModel

import "gofiber/app/models/model/types"

const (
	// CountryStatusActive 开启
	CountryStatusActive = 10

	// CountryStatusDisable 禁用
	CountryStatusDisable = -1
)

// Country 系统国家
type Country struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	Name    string `gorm:"type:varchar(60) not null;comment:名称" json:"name"`
	Alias   string `gorm:"type:varchar(60) not null;comment:别名" json:"alias"`
	Icon    string `gorm:"type:varchar(60) not null;comment:图标" json:"icon"`
	Iso1    string `gorm:"type:varchar(60) not null;comment:二位代码" json:"iso1"`
	Sort    int    `gorm:"type:tinyint not null;default:99;comment:排序" json:"sort"`
	Code    string `gorm:"type:varchar(50) not null;comment:区号" json:"code"`
	Status  int    `gorm:"type:smallint not null;default:10;comment:状态 -1禁用 10开启" json:"status"`
	Data    string `gorm:"type:text;comment:数据" json:"data"`
}

// SystemCountryInfo 管理系统国家信息
type SystemCountryInfo struct {
	ID   uint   `json:"id"`                       // ID
	Icon string `json:"icon"`                     // 国家图标
	Name string `gorm:"column:alias" json:"name"` // 国家别名
	Iso1 string `json:"iso1"`                     //	二位代码
	Code string `json:"code"`                     //	国家区号
}
