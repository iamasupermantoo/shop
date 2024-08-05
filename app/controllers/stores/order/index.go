package userOrder

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	Status     int                `json:"status"`     // 订单状态
	Pagination *scopes.Pagination `json:"pagination"` // 分页数据
}

// userStoreOrder 用户订单列表
type userStoreOrder struct {
	shopsModel.ProductStoreOrder
	StoreInfo shopsModel.Store `json:"storeInfo" gorm:"foreignKey:ID;references:StoreId"`
	OrderList []*productOrder  `json:"orderList" gorm:"foreignKey:OrderSn;references:OrderSn"`
}

func (userStoreOrder) TableName() string {
	return "product_store_order"
}

type productOrder struct {
	productsModel.ProductOrder
	ProductInfo productsModel.Product `json:"productInfo" gorm:"foreignKey:ID;references:ProductId"`
}

// Index 订单列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*userStoreOrder, 0)}
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Where("user_id = ?", ctx.UserId).
		Preload("StoreInfo").Preload("OrderList").Preload("OrderList.ProductInfo").
		Scopes(scopes.NewScopes().
			Eq("status", params.Status).
			Scopes()).
		Scopes(params.Pagination.Scopes()).
		Count(&data.Count).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
