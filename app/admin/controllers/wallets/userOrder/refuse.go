package userOrder

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type RefuseParams struct {
	Id   uint   `validate:"required" json:"id"`   //	ID
	Data string `validate:"required" json:"data"` //	拒绝理由
}

// Refuse 拒绝钱包订单
func Refuse(ctx *context.CustomCtx, params *RefuseParams) error {
	orderInfo := &walletsModel.WalletUserOrder{}
	result := database.Db.Model(orderInfo).Where("id = ?", params.Id).
		Where("status = ?", walletsModel.WalletUserOrderStatusActive).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(orderInfo)
	if result.Error != nil || orderInfo.ID == 0 {
		return ctx.ErrorJson("找不到可操作的订单 => " + result.Error.Error())
	}

	//	操作拒绝
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		result = tx.Model(&walletsModel.WalletUserOrder{}).Where("id = ?", orderInfo.ID).Updates(map[string]interface{}{
			"status": walletsModel.WalletUserOrderStatusRefuse,
			"data":   params.Data,
		})
		if result.Error != nil {
			return ctx.ErrorJson("订单修改失败")
		}

		userInfo := &usersModel.User{}
		result = database.Db.Model(userInfo).Where("id = ?", orderInfo.UserId).Find(userInfo)
		if result.Error != nil || userInfo.ID == 0 {
			return ctx.ErrorJson("找不到用户信息")
		}

		var err error
		switch orderInfo.Type {
		case walletsModel.WalletUserOrderTypeWithdraw:

			err = walletsService.NewUserWallet(tx, userInfo, nil).ChangeUserBalance(walletsModel.WalletUserBillTypeWithdrawRefuse, orderInfo.ID, orderInfo.Money)
		case walletsModel.WalletUserOrderTypeAssetsWithdraw:
			assetsInfo := &walletsModel.WalletAssets{}
			result = database.Db.Model(assetsInfo).Where("id = ?", orderInfo.AssetsId).Find(assetsInfo)
			if result.Error != nil || assetsInfo.ID == 0 {
				return ctx.ErrorJson("找不到资产信息")
			}
			err = walletsService.NewUserWallet(tx, userInfo, assetsInfo).ChangeUserBalance(walletsModel.WalletUserBillTypeAssetsWithdrawRefuse, orderInfo.ID, orderInfo.Money)
		}

		return err
	})

	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}
	return ctx.SuccessJsonOK()
}
