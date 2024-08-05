package storeRefund

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

type AgreeParams struct {
	OrderId  uint `json:"orderId" validate:"required"`  //	产品订单ID
	RefundId uint `json:"refundId" validate:"required"` //	售后ID
}

// Agree 售后订单同意
func Agree(ctx *context.CustomCtx, params *AgreeParams) error {
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

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		// 更新用户售后信息
		result = tx.Model(&shopsModel.StoreRefund{}).Where("id = ?", refundInfo.ID).Updates(&shopsModel.StoreRefund{
			Status: shopsModel.StoreRefundStatusComplete,
		})
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}

		// 更新当前订单
		result = tx.Model(&productsModel.ProductOrder{}).Where("id = ?", orderInfo.ID).Update("status", productsModel.ProductOrderStatusComplete)
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}

		// 查询用户信息
		userInfo := &usersModel.User{}
		result = database.Db.Model(userInfo).Where("id = ?", refundInfo.UserId).Find(userInfo)
		if result.Error != nil || userInfo.ID == 0 {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}

		// 返回用户余额
		return walletsService.NewUserWallet(tx, userInfo, nil).ChangeUserBalance(walletsModel.WalletUserBillTypeRefundProduct, orderInfo.ID, orderInfo.FinalMoney)
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	return ctx.SuccessJsonOK()
}
