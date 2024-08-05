package productsModel

import (
	"gofiber/app/models/model/types"
)

const (
	// ProductAttrsKeyStatusActivate 激活状态
	ProductAttrsKeyStatusActivate = 10
)

// ProductAttrsKey 产品属性名称
type ProductAttrsKey struct {
	types.GormModel
	ProductId uint   `json:"productId" gorm:"type:int unsigned not null;default:0;comment:产品ID"`
	Name      string `json:"name" gorm:"type:varchar(191) not null;default:'';comment:属性名称"`
	Type      int    `json:"type" gorm:"type:tinyint not null;default:1;comment:类型 1商品属性"`
	Data      string `json:"data" gorm:"type:varchar(255) not null;default:'';comment:数据"`
	Status    int    `json:"status" gorm:"type:tinyint not null;default:10;comment:状态 -1禁用｜10启用"`
}

// ProductAttrsKeyVal 产品属性Key
type ProductAttrsKeyVal struct {
	ProductAttrsKey
	Values []*ProductAttrsVal `json:"values" gorm:"foreignKey:KeyId;references:ID"`
}

func (_ProductAttrsKey *ProductAttrsKey) TableName() string {
	return "product_attrs_key"
}
