package productsModel

import (
	"gofiber/app/models/model/types"
)

const (
	// ProductAttrsValStatusActivate 激活状态
	ProductAttrsValStatusActivate = 10
)

// ProductAttrsVal 产品属性值
type ProductAttrsVal struct {
	types.GormModel
	KeyId  uint   `json:"keyId" gorm:"type:int unsigned not null;default:0;comment:属性名称ID"`
	Name   string `json:"name" gorm:"type:varchar(191) not null;default:'';comment:属性值名称"`
	Data   string `json:"data" gorm:"type:varchar(255) not null;default:'';comment:数据"`
	Status int    `json:"status" gorm:"type:tinyint not null;default:10;comment:状态 -1禁用｜10启用"`
}
