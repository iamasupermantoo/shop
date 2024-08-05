package userOrder

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type CancelParams struct {
	ID uint `json:"id" validate:"required"` //	店铺订单ID
}

// Cancel 订单取消
func Cancel(ctx *context.CustomCtx, params *CancelParams) error {
	// 获取订单信息
	storeOrderInfo := &shopsModel.ProductStoreOrder{}
	result := database.Db.Model(storeOrderInfo).Where("id = ?", params.ID).
		Where("user_id = ?", ctx.UserId).
		Where("status IN ?", []int{shopsModel.ProductStoreOrderStatusPending, shopsModel.ProductStoreOrderStatusShipping}).
		Find(storeOrderInfo)
	if result.Error != nil || storeOrderInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	if storeOrderInfo.Status == shopsModel.ProductStoreOrderStatusPending || storeOrderInfo.Status == shopsModel.ProductStoreOrderStatusShipping {
		err := database.Db.Transaction(func(tx *gorm.DB) error {
			result = tx.Model(&shopsModel.ProductStoreOrder{}).Where("id = ?", storeOrderInfo.ID).Update("status", shopsModel.ProductStoreOrderStatusDisable)
			if result.Error != nil {
				return ctx.ErrorJsonTranslate("abnormalOperation")
			}

			// 订单修改状态取消
			orderList := make([]*productsModel.ProductOrder, 0)
			database.Db.Model(&productsModel.ProductOrder{}).Where("store_order_id = ?", storeOrderInfo.ID).Find(&orderList)
			for _, orderInfo := range orderList {
				result = tx.Model(&productsModel.ProductOrder{}).Where("id = ?", orderInfo.ID).Update("status", productsModel.ProductOrderStatusDisable)
				if result.Error != nil {
					return ctx.ErrorJsonTranslate("abnormalOperation")
				}
			}

			// 如果已经支付了, 那么退还用户金额
			if storeOrderInfo.Status == shopsModel.ProductStoreOrderStatusShipping {
				userInfo := &usersModel.User{}
				result = database.Db.Model(userInfo).Where("id = ?", storeOrderInfo.UserId).Find(userInfo)
				if result.Error != nil || userInfo.ID == 0 {
					return ctx.ErrorJsonTranslate("abnormalOperation")
				}
				return walletsService.NewUserWallet(tx, userInfo, nil).ChangeUserBalance(walletsModel.WalletUserBillTypeRefundProduct, storeOrderInfo.ID, storeOrderInfo.FinalMoney)
			}
			return nil
		})
		if err != nil {
			return ctx.ErrorJsonTranslate(err.Error())
		}

	}

	return ctx.SuccessJsonOK()
}
