package refund

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type CreateParams struct {
	ID     uint     `json:"id" validate:"required"`     // 订单ID
	Name   string   `json:"name" validate:"required"`   //	申请理由
	Images []string `json:"images" validate:"required"` //	凭证图片
}

// Create 申请售后
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	orderInfo := &productsModel.ProductOrder{}
	result := database.Db.Model(orderInfo).Where("id = ?", params.ID).
		Where("status > ?", productsModel.ProductOrderStatusDisable).
		Where("user_id = ?", ctx.UserId).Find(orderInfo)
	if result.Error != nil || orderInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	// 售后只能有一个
	refundInfo := &shopsModel.StoreRefund{}
	database.Db.Model(refundInfo).Where("order_id = ?", params.ID).Where("user_id = ?", ctx.UserId).Find(refundInfo)
	if refundInfo.ID == 0 {
		result = database.Db.Create(&shopsModel.StoreRefund{
			AdminId: ctx.AdminId, UserId: ctx.UserId, OrderId: orderInfo.ID, ProductId: orderInfo.ProductId,
			Name: params.Name, Images: params.Images, Money: orderInfo.FinalMoney, StoreId: orderInfo.StoreId,
		})
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}
	}
	return ctx.SuccessJsonOK()
}
