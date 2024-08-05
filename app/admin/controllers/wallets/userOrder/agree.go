package userOrder

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type AgreeParams struct {
	Id uint `validate:"required" json:"id"` //	ID
}

// Agree 同意钱包订单操作
func Agree(ctx *context.CustomCtx, params *AgreeParams) error {
	orderInfo := &walletUserOrder{}
	result := database.Db.Model(&walletsModel.WalletUserOrder{}).Preload("PaymentInfo").
		Where("id = ?", params.Id).
		Where("status = ?", walletsModel.WalletUserOrderStatusActive).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(orderInfo)
	if result.Error != nil || orderInfo.ID == 0 {
		return ctx.ErrorJson("找不到可操作的订单")
	}

	//	操作同意
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		// 同意当前订单
		result = tx.Model(&walletsModel.WalletUserOrder{}).Where("id = ?", orderInfo.ID).Update("status", walletsModel.WalletUserOrderStatusComplete)
		if result.Error != nil {
			return ctx.ErrorJson("订单修改失败")
		}

		userInfo := &usersModel.User{}
		result = database.Db.Model(userInfo).Where("id = ?", orderInfo.UserId).Find(userInfo)
		if result.Error != nil {
			return ctx.ErrorJson("用户不存在")
		}
		var err error
		switch orderInfo.Type {
		case walletsModel.WalletUserOrderTypeDeposit:
			rateMoney := orderInfo.Money * orderInfo.PaymentInfo.Rate
			err = walletsService.NewUserWallet(tx, userInfo, nil).ChangeUserBalance(walletsModel.WalletUserBillTypeDeposit, orderInfo.ID, rateMoney)
		case walletsModel.WalletUserOrderTypeAssetsDeposit:
			assetsInfo := &walletsModel.WalletAssets{}
			result = database.Db.Model(assetsInfo).Where("id = ?", orderInfo.AssetsId).Find(assetsInfo)
			if result.Error != nil || assetsInfo.ID == 0 {
				return ctx.ErrorJson("找不到资产信息")
			}
			err = walletsService.NewUserWallet(tx, userInfo, assetsInfo).ChangeUserAssets(walletsModel.WalletUserBillTypeAssetsDeposit, orderInfo.ID, orderInfo.Money)
		}

		return err
	})

	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}
	return ctx.SuccessJsonOK()
}
