package productsModel

import (
	"gofiber/app/models/model/types"
)

const (
	// ProductAttrsSkuStatusActivate 激活状态
	ProductAttrsSkuStatusActivate = 10
)

// ProductAttrsSku 产品属性SKU
type ProductAttrsSku struct {
	types.GormModel
	ParentId  uint    `json:"parentId" gorm:"type:int unsigned not null;comment:父级ID"`
	ProductId uint    `json:"productId" gorm:"type:int unsigned not null;default:0;comment:产品ID"`
	Vals      string  `json:"vals" gorm:"type:varchar(255) not null;default:'';comment:属性值ID,用逗号分隔"`
	Name      string  `json:"name" gorm:"type:varchar(512) not null;default:'';comment:SKU名称"`
	Image     string  `json:"image" gorm:"type:varchar(255) not null;default:'';comment:商品图片"`
	Stock     uint    `json:"stock" gorm:"type:int unsigned not null;default:1000;comment:库存量"`
	Sales     uint    `json:"sales" gorm:"type:int unsigned not null;default:0;comment:销售量"`
	Money     float64 `json:"money" gorm:"type:decimal(12, 2) not null;default:100;comment:标价"`
	Discount  float64 `json:"discount" gorm:"type:decimal(8,4) not null;default:0;comment:折扣"`
	Data      string  `json:"data" gorm:"type:varchar(255) not null;default:'';comment:数据"`
	Status    int     `json:"status" gorm:"type:tinyint not null;default:10;comment:状态 -1下架｜10上架"`
}

// ProductAttrsSkuList 产品属性SKU列表
type ProductAttrsSkuList struct {
	Ids  []string //	属性值ID列表
	Name []string //	属性值名称
}

// GetTotalPrice 获取总价
func (_ProductAttrsSku *ProductAttrsSku) GetTotalPrice(nums float64) float64 {
	return _ProductAttrsSku.Money * nums
}

// GetFinalPrice 获取最终价
func (_ProductAttrsSku *ProductAttrsSku) GetFinalPrice(nums float64) float64 {
	return _ProductAttrsSku.GetTotalPrice(nums) - _ProductAttrsSku.GetTotalPrice(nums)*_ProductAttrsSku.Discount
}

// GetEarning 获取收益
func (_ProductAttrsSku *ProductAttrsSku) GetEarning(nums, increase float64) float64 {
	return _ProductAttrsSku.GetFinalPrice(nums) - _ProductAttrsSku.GetFinalPrice(nums)*(increase/100)
}
