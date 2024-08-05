package storeOrder

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	Status     int                `json:"status"`     // 订单状态
	Pagination *scopes.Pagination `json:"pagination"` // 分页数据
}

// StoreUserOrder 店铺订单列表
type StoreUserOrder struct {
	shopsModel.ProductStoreOrder
	AddressInfo shopsModel.ShippingAddress `json:"addressInfo" gorm:"-"`
	UserInfo    usersModel.UserInfo        `json:"userInfo" gorm:"foreignKey:UserId"`
	OrderList   []*productOrder            `json:"orderList" gorm:"foreignKey:OrderSn;references:OrderSn"`
}

func (*StoreUserOrder) TableName() string {
	return "product_store_order"
}

type productOrder struct {
	productsModel.ProductOrder
	ProductInfo productsModel.Product `json:"productInfo" gorm:"foreignKey:ProductId"`
}

// Index 订单列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*StoreUserOrder, 0)}
	storeInfo := &shopsModel.Store{}
	database.Db.Model(storeInfo).Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.StoreStatusActivate).Find(storeInfo)

	database.Db.Model(&shopsModel.ProductStoreOrder{}).Where("store_id = ?", storeInfo.ID).
		Preload("UserInfo").Preload("OrderList").Preload("OrderList.ProductInfo").
		Scopes(scopes.NewScopes().
			Eq("status", params.Status).Scopes()).
		Scopes(params.Pagination.Scopes()).
		Count(&data.Count).
		Find(&data.Items)
	return ctx.SuccessJson(data)
}
