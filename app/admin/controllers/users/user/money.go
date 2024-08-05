package user

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type MoneyParams struct {
	UserName string  `validate:"required" json:"userName"` //	用户账户
	Money    float64 `validate:"required" json:"money"`    //	金额
}

func Money(ctx *context.CustomCtx, params *MoneyParams) error {
	userInfo := &usersModel.User{}
	result := database.Db.Model(userInfo).Where("user_name = ?", params.UserName).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJson("查询用户失败")
	}

	resultBillType := walletsModel.WalletUserBillTypeSystemDeposit
	resultMoney := params.Money
	if params.Money < 0 {
		resultBillType = walletsModel.WalletUserBillTypeSystemWithdraw
		resultMoney = -params.Money
	}

	// 事务操作用户资产金额
	return ctx.IsErrorJson(database.Db.Transaction(func(tx *gorm.DB) error {
		return walletsService.NewUserWallet(tx, userInfo, nil).ChangeUserBalance(resultBillType, 0, resultMoney)
	}))
}
