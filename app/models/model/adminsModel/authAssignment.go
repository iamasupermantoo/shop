package adminsModel

import "gofiber/app/models/model/types"

// AuthAssignment 权限分配表
type AuthAssignment struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;index:idx_admin_name;comment:管理ID" json:"adminId"`
	Name    string `gorm:"type:varchar(50) not null;index:idx_admin_name;comment:名称" json:"name"`
}
