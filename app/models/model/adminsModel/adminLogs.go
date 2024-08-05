package adminsModel

import "gofiber/app/models/model/types"

// AdminLogs 管理日志记录
type AdminLogs struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	Ip      string `gorm:"type:varchar(120) not null;comment:IP地址" json:"ip"`
	Headers string `gorm:"type:text;comment:头信息" json:"headers"`
	Name    string `gorm:"type:varchar(120) not null;comment:名称" json:"name"`
	Route   string `gorm:"type:varchar(60) not null;comment:路由" json:"route"`
	Body    string `gorm:"type:text;comment:参数" json:"body"`
	Data    string `gorm:"type:text;comment:数据" json:"data"`
}
