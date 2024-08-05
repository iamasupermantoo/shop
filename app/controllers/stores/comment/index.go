package comment

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

// Index 评论列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: []*UserComment{}}
	productOrderIds := make([]int, 0)
	database.Db.Model(&productsModel.ProductOrder{}).Where("user_id = ?", ctx.UserId).Pluck("id", &productOrderIds)
	if len(productOrderIds) == 0 {
		return ctx.SuccessJson(data)
	}
	database.Db.Model(&shopsModel.StoreComment{}).Preload("UserInfo").
		Preload("StoreInfo").
		Preload("ProductInfo").
		Preload("OrderInfo").
		Where("order_id IN ?", productOrderIds).
		Scopes(scopes.NewScopes().
			Eq("status", params.Status).Scopes()).
		Scopes(params.Pagination.Scopes()).
		Count(&data.Count).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}

type UserComment struct {
	shopsModel.StoreComment
	UserInfo    usersModel.User            `json:"userInfo" gorm:"foreignKey:UserId"`
	StoreInfo   shopsModel.Store           `json:"storeInfo" gorm:"foreignKey:StoreId"`
	ProductInfo productsModel.Product      `json:"productInfo" gorm:"foreignKey:ProductId"`
	OrderInfo   productsModel.ProductOrder `json:"orderInfo" gorm:"foreignKey:OrderId"`
}
