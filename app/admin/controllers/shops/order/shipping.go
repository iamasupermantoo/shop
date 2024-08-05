package shopsOrder

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/shopsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type ShippingParams struct {
	ID int `json:"id" validate:"required"` //	店铺订单ID
}

// Shipping 发货
func Shipping(ctx *context.CustomCtx, params *ShippingParams) error {
	// 订单信息
	storeOrderInfo := shopsModel.ProductStoreOrder{}
	result := database.Db.Model(&storeOrderInfo).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Where("id = ?", params.ID).
		Find(&storeOrderInfo)
	if result.Error != nil || storeOrderInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	var storeUserId uint
	database.Db.Raw("select s.user_id from product_store_order as pso left join store as s on pso.store_id = s.id where pso.id = ?", storeOrderInfo.ID).Scan(&storeUserId)
	userInfo := usersModel.User{}
	result = database.Db.Model(&userInfo).
		Where("id = ?", storeUserId).
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
