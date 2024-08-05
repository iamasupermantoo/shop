package userOrder

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type DeleteParams struct {
	ID uint `json:"id" validate:"required"` //	店铺订单ID
}

// Delete 订单删除
func Delete(ctx *context.CustomCtx, params *DeleteParams) error {
	// 获取订单信息
	storeOrderInfo := &shopsModel.ProductStoreOrder{}
	result := database.Db.Model(storeOrderInfo).Where("id = ?", params.ID).Where("user_id = ?", ctx.UserId).Find(storeOrderInfo)
	if result.Error != nil || storeOrderInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	// 取消订单｜完成订单 可以删除
	if storeOrderInfo.Status == shopsModel.ProductStoreOrderStatusDisable || storeOrderInfo.Status == shopsModel.ProductStoreOrderStatusComplete {
		err := database.Db.Transaction(func(tx *gorm.DB) error {
			result = tx.Where("id = ?", storeOrderInfo.ID).Delete(&shopsModel.ProductStoreOrder{})
			if result.Error != nil {
				return ctx.ErrorJsonTranslate("abnormalOperation")
			}

			orderList := make([]*productsModel.ProductOrder, 0)
			database.Db.Model(&productsModel.ProductOrder{}).Where("store_order_id = ?", storeOrderInfo.ID).Find(&orderList)
			for _, orderInfo := range orderList {
				result = tx.Where("id = ?", orderInfo.ID).Delete(&productsModel.ProductOrder{})
				if result.Error != nil {
					return ctx.ErrorJsonTranslate("abnormalOperation")
				}
			}
			return nil
		})

		if err != nil {
			return ctx.ErrorJsonTranslate(err.Error())
		}
	}

	return ctx.SuccessJsonOK()
}
