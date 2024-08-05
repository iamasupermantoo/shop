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

type PayParams struct {
	ID       uint `json:"id"`       // 订单Id
	AssetsId uint `json:"assetsId"` // 资产ID
}

// Pay 支付订单
func Pay(ctx *context.CustomCtx, params *PayParams) error {
	userInfo := usersModel.User{}
	result := database.Db.Where("id = ?", ctx.UserId).Find(&userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	storeOrderInfo := shopsModel.ProductStoreOrder{}
	database.Db.Where("id = ?", params.ID).
		Where("user_id = ?", ctx.UserId).
		Where("status = ?", shopsModel.ProductStoreOrderStatusPending).
		Find(storeOrderInfo)
	if storeOrderInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		if result = tx.Where("id = ?", params.ID).Updates(&shopsModel.ProductStoreOrder{
			AssetsId: params.AssetsId,
			Status:   shopsModel.ProductStoreOrderStatusShipping,
		}); result.Error != nil {
			return result.Error
		}

		if result = tx.Where("id = ?", storeOrderInfo.StoreId).Updates(&productsModel.ProductOrder{
			AssetsId: params.AssetsId,
			Status:   productsModel.ProductOrderStatusShipping,
		}); result.Error != nil {
			return result.Error
		}

		return walletsService.NewUserWallet(tx, &userInfo, nil).ChangeUserBalance(walletsModel.WalletUserBillTypeBuyProduct, storeOrderInfo.ID, storeOrderInfo.FinalMoney)
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	return nil
}
