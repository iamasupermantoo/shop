package adminsModel

import "gofiber/app/models/model/types"

const (
	//	AuthItemTypeRole 权限角色
	AuthItemTypeRole = 1

	//	AuthItemTypeRoute 权限路由
	AuthItemTypeRoute = 2

	//	AuthItemTypeName 路由名称类型
	AuthItemTypeName = 3

	// AuthRoleSuperManage 超级管理员
	AuthRoleSuperManage = "超级管理员"

	// AuthRoleMerchantManage 商户管理员
	AuthRoleMerchantManage = "商户管理员"

	// AuthRoleAgentManage 代理管理员
	AuthRoleAgentManage = "代理管理员"
)

// AuthItem 权限目录
type AuthItem struct {
	types.GormModel
	Name string `gorm:"index;type:varchar(50) not null;comment:名称" json:"name"`
	Type int    `gorm:"type:tinyint not null;comment:类型 1权限角色 2路由权限 3路由名称" json:"type"`
	Desc string `gorm:"type:varchar(255) not null;comment:详情" json:"desc"`
	Rule string `gorm:"type:varchar(255) not null;comment:规则" json:"rule"`
	Data string `gorm:"type:varchar(255) not null;comment:数据" json:"data"`
}
