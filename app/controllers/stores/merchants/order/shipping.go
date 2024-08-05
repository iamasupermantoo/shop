package storeOrder

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/shopsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type ShippingParams struct {
	ID uint `json:"id" validate:"required"` //	店铺订单ID
}

// Shipping 发货
func Shipping(ctx *context.CustomCtx, params *ShippingParams) error {
	storeInfo := shopsModel.Store{}
	result := database.Db.Model(&storeInfo).
		Where("user_id = ?", ctx.UserId).
		Where("status = ?", shopsModel.StoreStatusActivate).
		Find(&storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	// 订单信息
	storeOrderInfo := shopsModel.ProductStoreOrder{}
	result = database.Db.Model(&storeOrderInfo).
		Where("store_id = ?", storeInfo.ID).
		Where("id = ?", params.ID).
		Find(&storeOrderInfo)
	if result.Error != nil || storeOrderInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	userInfo := usersModel.User{}
	result = database.Db.Model(&userInfo).
		Where("id = ?", ctx.UserId).
		Find(&userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		return shopsService.NewStoreOrder(tx).Shipping(userInfo, storeOrderInfo)
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	return ctx.SuccessJsonOK()
}
