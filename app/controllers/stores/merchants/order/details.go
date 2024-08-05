package storeOrder

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type DetailsParams struct {
	ID uint `json:"id" validate:"required"` //	店铺订单ID
}

// Details 店铺订单详情
func Details(ctx *context.CustomCtx, params *DetailsParams) error {
	orderInfo := &StoreUserOrder{}
	storeInfo := &shopsModel.Store{}
	database.Db.Model(storeInfo).Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.StoreStatusActivate).Find(storeInfo)
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Where("id = ?", params.ID).Where("store_id = ?", storeInfo.ID).
		Preload("UserInfo").Preload("OrderList").Preload("OrderList.ProductInfo").
		Find(&orderInfo)
	return ctx.SuccessJson(orderInfo)
}
