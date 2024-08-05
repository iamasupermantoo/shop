package refund

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type RefuseParams struct {
	OrderId  uint   `json:"orderId" validate:"required"`  //	订单ID
	RefundId uint   `json:"refundId" validate:"required"` //	售后ID
	Data     string `json:"data" validate:"required"`     //	拒绝理由
}

// Refuse 售后订单拒绝
func Refuse(ctx *context.CustomCtx, params *RefuseParams) error {
	storeInfo := &shopsModel.Store{}
	result := database.Db.Model(storeInfo).Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.StoreStatusActivate).Find(storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	// 查询订单
	orderInfo := &productsModel.ProductOrder{}
	result = database.Db.Model(orderInfo).Where("id = ?", params.OrderId).Where("store_id = ?", storeInfo.ID).Find(orderInfo)
	if result.Error != nil || orderInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	// 查询售后ID
	refundInfo := &shopsModel.StoreRefund{}
	result = database.Db.Model(refundInfo).Where("id = ?", params.RefundId).Where("order_id = ?", params.OrderId).Find(refundInfo)
	if result.Error != nil || refundInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	result = database.Db.Model(&shopsModel.StoreRefund{}).Where("id = ?", refundInfo.ID).Updates(&shopsModel.StoreRefund{
		Status: shopsModel.StoreRefundStatusRefuse, Data: params.Data,
	})
	if result.Error != nil {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	return ctx.SuccessJsonOK()
}
