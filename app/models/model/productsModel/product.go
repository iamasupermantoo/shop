package productsModel

import "gofiber/app/models/model/types"

const (
	// ProductStatusActive 激活
	ProductStatusActive = 10

	// ProductStatusDisable 禁用
	ProductStatusDisable = -1

	// ProductTypeDefault 店铺商品类型
	ProductTypeDefault = 1

	// ProductTypeWholesale  批发商品类型
	ProductTypeWholesale = 2
)

// Product 产品表
type Product struct {
	types.GormModel
	AdminId    uint              `json:"adminId" gorm:"type:int unsigned not null;comment:管理员ID"`
	ParentId   uint              `json:"parentId" gorm:"type:int unsigned not null;comment:父级ID"`
	StoreId    uint              `json:"storeId" gorm:"type:int unsigned not null;comment:店铺ID"`
	CategoryId uint              `json:"categoryId" gorm:"type:int unsigned not null;comment:类目ID"`
	AssetsId   uint              `json:"assetsId" gorm:"type:int unsigned not null;comment:资产ID"`
	Name       string            `json:"name" gorm:"type:varchar(1024) not null;comment:标题"`
	Images     types.GormStrings `json:"images" gorm:"type:varchar(2048) not null;comment:图标"`
	Money      float64           `json:"money" gorm:"type:decimal(20,6) not null;comment:金额"`
	Discount   float64           `json:"discount" gorm:"type:decimal(8,4) not null;comment:折扣"`
	Type       int               `json:"type" gorm:"type:tinyint not null;default:1;comment:类型1默认类型"`
	Sort       int               `json:"sort" gorm:"type:tinyint not null;default:99;comment:排序"`
	Sales      int               `json:"sales" gorm:"type:int unsigned not null;default:0;comment:销售量"`
	Status     int               `json:"status" gorm:"type:tinyint not null; default:10;comment:状态-1禁用 10启用"`
	Data       string            `json:"data" gorm:"type:text;comment:数据"`
	Desc       string            `json:"desc" gorm:"type:text;comment:描述"`
}
