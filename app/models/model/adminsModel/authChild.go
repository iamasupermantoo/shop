package adminsModel

import "gofiber/app/models/model/types"

const (
	// AuthChildTypeRoleParentRole 角色对应角色
	AuthChildTypeRoleParentRole = 1

	// AuthChildTypeRouteNameRoute 路由名称对应路由
	AuthChildTypeRouteNameRoute = 2

	// AuthChildTypeRoleRouteName 角色对应路由名称
	AuthChildTypeRoleRouteName = 3
)

// AuthChild 权限目录子级表
type AuthChild struct {
	types.GormModel
	Parent string `gorm:"type:varchar(50) not null;index:idx_parent_child;comment:父级" json:"parent"`
	Child  string `gorm:"type:varchar(50) not null;index:idx_parent_child;comment:子级" json:"child"`
	Type   int    `gorm:"type:tinyint not null;comment:类型 2路由名称对应路由 3角色对应路由名" json:"type"`
}
