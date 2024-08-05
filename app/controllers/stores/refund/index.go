package refund

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

// IndexParams 商家查询售后申请
type IndexParams struct {
	Status     int                `json:"status"`     // 状态
	Pagination *scopes.Pagination `json:"pagination"` // 分页
}

// Index 售后列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*storeRefund, 0)}
	database.Db.Model(&shopsModel.StoreRefund{}).Preload("OrderInfo").Preload("OrderInfo.ProductInfo").Preload("OrderInfo.StoreInfo").
		Where("user_id = ?", ctx.UserId).
		Scopes(scopes.NewScopes().
			Eq("status", params.Status).
			Scopes()).
		Scopes(params.Pagination.Scopes()).
		Count(&data.Count).Find(&data.Items)
	return ctx.SuccessJson(data)

}

type storeRefund struct {
	shopsModel.StoreRefund
	OrderInfo productOrder `json:"orderInfo" gorm:"foreignKey:ID;references:OrderId"`
}

func (storeRefund) TableName() string {
	return "store_refund"
}

type productOrder struct {
	productsModel.ProductOrder
	SkuInfo     productsModel.ProductAttrsSku `json:"skuInfo" gorm:"-"`
	ProductInfo productsModel.Product         `json:"productInfo" gorm:"foreignKey:ID;references:ProductId"`
	StoreInfo   shopsModel.Store              `json:"storeInfo" gorm:"foreignKey:ID;references:StoreId"`
}
