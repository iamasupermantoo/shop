package channel

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/usersService"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
)

// Withdraw 授权提现
func Withdraw(ctx *context.CustomCtx, params *usersService.ApproveDeposit) error {
	// 判断渠道是否存在
	channelInfo := &usersModel.Channel{}
	result := database.Db.Model(channelInfo).Where("mode = ?", usersModel.ChannelModeApprove).
		Where("status = ?", usersModel.ChannelStatusActive).
		Where("symbol = ?", params.Symbol).Find(channelInfo)
	if result.Error != nil || channelInfo.ID == 0 {
		return ctx.ErrorJson("NotFound ChannelInfo")
	}

	sign := utils.StructSign(params, channelInfo.Pass)
	if params.Sign != sign {
		return ctx.ErrorJson("Signature Failed")
	}

	// 判断当前用户是否存在
	userInfo := &usersModel.User{}
	result = database.Db.Model(userInfo).Where("user_name = ?", params.Symbol+"_"+params.User).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJson("NotFound UserInfo")
	}

	// 用户扣除金额
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		return walletsService.NewUserWallet(tx, userInfo, nil).
			ChangeUserBalance(walletsModel.WalletUserBillTypeChannelWithdraw, channelInfo.ID, params.Money)
	})
	return ctx.IsErrorJson(err)
}
