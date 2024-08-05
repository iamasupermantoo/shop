package storeComment

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type ProductParams struct {
	Id         uint               `json:"id"`         // 产品Id
	Status     int                `json:"status"`     // 订单状态
	Pagination *scopes.Pagination `json:"pagination"` // 分页数据
}

// Product  产品评论列表
func Product(ctx *context.CustomCtx, params *ProductParams) error {
	data := &context.IndexData{Items: []*ProductComment{}}
	database.Db.Model(&shopsModel.StoreComment{}).
		Where("product_id = ?", params.Id).
		Preload("StoreInfo").
		Preload("ProductInfo").
		Preload("UserInfo").
		Preload("OrderInfo").
		Scopes(scopes.NewScopes().
			Eq("status", params.Status).
			Scopes()).
		Scopes(params.Pagination.Scopes()).
		Count(&data.Count).Find(&data.Items)

	return ctx.SuccessJson(data)
}

type ProductComment struct {
	shopsModel.StoreComment
	SkuInfo     productsModel.ProductAttrsSku `json:"skuInfo" gorm:"-"`
	StoreInfo   shopsModel.Store              `json:"storeInfo" gorm:"foreignKey:ID;references:StoreId"`
	UserInfo    usersModel.User               `json:"userInfo" gorm:"foreignKey:ID;references:UserId"`
	ProductInfo productsModel.Product         `json:"productInfo" gorm:"foreignKey:ID;references:ProductId"`
	OrderInfo   productsModel.ProductOrder    `json:"orderInfo" gorm:"foreignKey:ID;references:OrderId"`
}

func (ProductComment) TableName() string {
	return "store_comment"
}
