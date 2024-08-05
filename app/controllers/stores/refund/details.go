package refund

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type DetailsParams struct {
	ID uint `json:"id" validate:"required"` //	产品订单ID
}

type orderRefundInfo struct {
	productsModel.ProductOrder
	ProductInfo productsModel.Product  `json:"productInfo" gorm:"foreignKey:ID;references:ProductId"`
	RefundInfo  shopsModel.StoreRefund `json:"refundInfo" gorm:"foreignKey:OrderId;references:ID"`
}

func (orderRefundInfo) TableName() string {
	return "product_order"
}

// Details 产品订单售后
func Details(ctx *context.CustomCtx, params *DetailsParams) error {
	refundInfo := orderRefundInfo{}
	database.Db.Model(&productsModel.ProductOrder{}).Preload("ProductInfo").Preload("RefundInfo").
		Where("id = ?", params.ID).
		Where("user_id = ?", ctx.UserId).
		Find(&refundInfo)

	return ctx.SuccessJson(refundInfo)
}
