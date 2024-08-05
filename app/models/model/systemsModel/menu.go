package systemsModel

import "gofiber/app/models/model/types"

const (
	// MenuStatusActive 菜单激活
	MenuStatusActive = 10
	// MenuStatusDisable 禁用
	MenuStatusDisable = -1
	// MenuTypeNavigation 导航菜单
	MenuTypeNavigation = 1
	// MenuTypeSetting 设置菜单
	MenuTypeSetting = 11
	// MenuTypeMore 更多菜单
	MenuTypeMore = 21
	// MenuTypeStore 商户菜单
	MenuTypeStore = 31
)

// Menu 前台菜单
type Menu struct {
	types.GormModel
	AdminId    uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	ParentId   uint   `gorm:"type:int unsigned not null;comment:父级ID" json:"parentId"`
	Name       string `gorm:"type:varchar(50) not null;comment:名称" json:"name"`
	Route      string `gorm:"type:varchar(50) not null;comment:路由" json:"route"`
	Sort       int    `gorm:"type:tinyint not null;default:99;comment:排序" json:"sort"`
	Icon       string `gorm:"type:varchar(255) not null;comment:图标" json:"icon"`
	ActiveIcon string `gorm:"type:varchar(255) not null;comment:选中图标" json:"activeIcon"`
	IsDesktop  int    `gorm:"type:tinyint(1) not null;default:1;comment:桌面显示" json:"isDesktop"`
	IsMobile   int    `gorm:"type:tinyint(1) not null;default:1;comment:手机显示" json:"isMobile"`
	Type       int    `gorm:"type:tinyint not null;default:1;comment:类型1导航菜单 11设置菜单 21更多菜单" json:"type"`
	Status     int    `gorm:"type:tinyint not null;default:10;comment:状态-1禁用 10开启" json:"status"`
	Data       string `gorm:"type:text;comment:数据" json:"data"`
}

// SystemMenuInfo 前端设置菜单
type SystemMenuInfo struct {
	Name       string            `json:"name"`       // 菜单名称
	Route      string            `json:"route"`      // 菜单路由
	Icon       string            `json:"icon"`       // 菜单图标
	ActiveIcon string            `json:"activeIcon"` // 菜单激活图标
	IsDesktop  int               `json:"isDesktop"`  // 桌面显示
	IsMobile   int               `json:"isMobile"`   // 手机显示
	Children   []*SystemMenuInfo `json:"children"`   // 子级
}
