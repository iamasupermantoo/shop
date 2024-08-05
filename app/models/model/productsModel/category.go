package productsModel

import "gofiber/app/models/model/types"

const (
	// CategoryStatusActive 分类激活
	CategoryStatusActive = 10
	// CategoryStatusDisable 分类禁用
	CategoryStatusDisable = -1

	// CategoryTypeDefault 默认分类
	CategoryTypeDefault = 1
)

// Category 分类表
type Category struct {
	types.GormModel
	ParentId uint   `json:"parentId" gorm:"type:int unsigned not null;comment:分类上级ID"`
	AdminId  uint   `json:"adminId" gorm:"type:int unsigned not null;comment:管理员ID"`
	Type     int    `json:"type" gorm:"type:tinyint not null;default:1;comment:类型1默认类型"`
	Name     string `json:"name" gorm:"type:varchar(60) not null;comment:标题"`
	Icon     string `json:"icon" gorm:"type:varchar(60) not null;comment:封面"`
	Sort     int    `json:"sort" gorm:"type:tinyint not null;default:99;comment:排序"`
	Status   int    `json:"status" gorm:"type:tinyint not null;default:10;comment:状态-1禁用 10启用"`
	Data     string `json:"data" gorm:"type:text;comment:数据"`
}
