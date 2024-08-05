package userAssets

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type MoneyParams struct {
	UserName string  `validate:"required" json:"userName"` //	用户账户
	AssetsId uint    `validate:"required" json:"assetsId"` //	ID
	Money    float64 `validate:"required" json:"money"`    //	金额
}

// Money 加减款
func Money(ctx *context.CustomCtx, params *MoneyParams) error {
	adminChildIds := ctx.GetAdminChildIds()
	userInfo := &usersModel.User{}
	result := database.Db.Model(userInfo).Where("user_name = ?", params.UserName).Where("admin_id IN ?", adminChildIds).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJson("查询用户失败")
	}

	assetsInfo := &walletsModel.WalletAssets{}
	result = database.Db.Model(assetsInfo).Where("id = ?", params.AssetsId).
		Where("admin_id IN ?", adminsService.NewAdminUser(ctx.Rds, userInfo.AdminId).GetRedisChildrenIds()).Find(assetsInfo)
	if result.Error != nil || assetsInfo.ID == 0 {
		return ctx.ErrorJson("查询资产名称失败")
	}

	resultBillType := walletsModel.WalletUserBillTypeSystemAssetsDeposit
	resultMoney := params.Money
	if params.Money < 0 {
		resultBillType = walletsModel.WalletUserBillTypeSystemAssetsWithdraw
		resultMoney = -params.Money
	}

	// 事务操作用户资产金额
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		return walletsService.NewUserWallet(tx, userInfo, assetsInfo).ChangeUserAssets(resultBillType, 0, resultMoney)
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	return ctx.SuccessJsonOK()
}
