package usersModel

import "gofiber/app/models/model/types"

const (
	// AccessTypeDefault 站点访问
	AccessTypeDefault = 1

	// AccessTypeStore 店铺访问
	AccessTypeStore = 2
)

// Access 前台访问表
type Access struct {
	types.GormModel
	AdminId  uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	UserId   uint   `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	SourceId uint   `gorm:"type:int unsigned not null;comment:来源ID" json:"sourceId"`
	Name     string `gorm:"type:varchar(120) not null;comment:名称" json:"name"`
	Type     int    `gorm:"type:tinyint not null;default:1;comment:类型" json:"type"`
	Ip       string `gorm:"type:varchar(120) not null;comment:IP地址" json:"ip"`
	Route    string `gorm:"type:varchar(60) not null;comment:路由" json:"route"`
	Headers  string `gorm:"type:text;comment:头信息" json:"headers"`
	Data     string `gorm:"type:text;comment:数据" json:"data"`
}
