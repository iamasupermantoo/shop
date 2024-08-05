package usersModel

import "gofiber/app/models/model/types"

// Setting 用户设置
type Setting struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	UserId  uint   `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	Name    string `gorm:"type:varchar(50) not null;comment:备注" json:"name"`
	Type    int    `gorm:"type:tinyint not null;default:1;comment:类型" json:"type"`
	Field   string `gorm:"type:varchar(50) not null;comment:建铭" json:"field"`
	Value   string `gorm:"type:text;comment:键值" json:"value"`
	Data    string `gorm:"type:text;comment:input配置" json:"data"`
}
