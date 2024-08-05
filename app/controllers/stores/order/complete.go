package userOrder

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/shopsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type CompleteParams struct {
	ID uint `json:"id" validate:"required"` //	订单ID
}

// Complete 订单完成
func Complete(ctx *context.CustomCtx, params *CompleteParams) error {
	storeOrderInfo := shopsModel.ProductStoreOrder{}
	result := database.Db.Model(&storeOrderInfo).
		Where("id = ?", params.ID).
		Where("status = ?", shopsModel.ProductStoreOrderStatusProgress).
		Where("user_id = ?", ctx.UserId).
		Find(&storeOrderInfo)
	if result.Error != nil {
		return ctx.ErrorJsonTranslate("findError")
	}

	storeInfo := shopsModel.Store{}
	result = database.Db.Model(&storeInfo).
		Where("id = ?", storeOrderInfo.StoreId).
		Find(&storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	storeUserInfo := usersModel.User{}
	result = database.Db.Model(&storeUserInfo).
		Where("id = ?", storeInfo.UserId).
		Find(&storeUserInfo)
	if result.Error != nil || storeUserInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		return shopsService.NewStoreOrder(tx).Complete(storeUserInfo, storeOrderInfo)
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	return ctx.SuccessJsonOK()
}
