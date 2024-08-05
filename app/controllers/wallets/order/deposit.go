package order

import (
	"gofiber/app/models/model/types"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/usersService"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/websocket"
	"gofiber/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type DepositParams struct {
	PaymentId uint    `json:"paymentId" validate:"required"`  // 支付ID
	Money     float64 `json:"money" validate:"required,gt=0"` // 金额
	Voucher   string  `json:"voucher"`                        // 支付凭证
}

// Deposit 钱包订单充值
func Deposit(ctx *context.CustomCtx, params *DepositParams) error {
	paymentInfo := walletsModel.WalletPayment{}
	result := database.Db.Model(&paymentInfo).Where("id = ?", params.PaymentId).
		Where("mode IN ?", []int{walletsModel.WalletPaymentModeDeposit, walletsModel.WalletPaymentModeAssetsDeposit}).
		Where("admin_id = ?", ctx.AdminSettingId).
		Where("status = ?", walletsModel.WalletPaymentStatusActive).
		Find(&paymentInfo)
	if result.Error != nil || paymentInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	// 如果是渠道充值, 那么调用渠道充值, 直接充值
	if paymentInfo.Type == walletsModel.WalletPaymentTypeChannel {
		return channelDeposit(ctx, &paymentInfo, params)
	}

	moneyRate := paymentInfo.Rate * params.Money
	adminSettingCache := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId)
	rangeMoney := adminSettingCache.GetRangeMoney("walletDepositAmountBetween")
	if moneyRate < rangeMoney.Min || moneyRate > rangeMoney.Max {
		return ctx.ErrorJsonTranslate("rangeMismatch", rangeMoney.Min, rangeMoney.Max)
	}

	// 开启了需要配置不能为空
	if params.Voucher == "" && paymentInfo.IsVoucher == types.ModelBoolTrue {
		return ctx.ErrorJsonTranslateMultiple("paymentVoucher", "notBeEmpty")
	}

	// 用户有存款订单则跳过，没有就创建订单
	if result = database.Db.Where("user_id = ?", ctx.UserId).
		Where("status = ?", walletsModel.WalletUserOrderStatusActive).
		Where("type = ?", paymentInfo.Mode).
		Find(&walletsModel.WalletUserOrder{}); result.RowsAffected != 0 {
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("limitExceeded")
		}
	}

	database.Db.Create(&walletsModel.WalletUserOrder{
		AdminId:  ctx.AdminId,
		UserId:   ctx.UserId,
		AssetsId: paymentInfo.AssetsId,
		Type:     paymentInfo.Mode,
		SourceId: paymentInfo.ID,
		OrderSn:  utils.NewRandom().OrderSn(), Money: params.Money, Voucher: params.Voucher,
	})

	// 管理通知
	userInfo := &usersModel.User{}
	database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Find(userInfo)
	websocket.NewAdminNotify(ctx.Rds, ctx.AdminId).Deposit(*userInfo, paymentInfo, params.Money)

	return ctx.SuccessJsonOK()
}

// channelDeposit 渠道充值
func channelDeposit(ctx *context.CustomCtx, paymentInfo *walletsModel.WalletPayment, params *DepositParams) error {
	// 只能支持余额
	if paymentInfo.Mode != walletsModel.WalletPaymentModeDeposit {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	var channelSymbol, channelIdStr string
	for _, datum := range paymentInfo.Data {
		if datum.Field == "realName" {
			channelSymbol = datum.Value
		}
		if datum.Field == "number" {
			channelIdStr = datum.Value
		}
	}
	channelId, _ := strconv.Atoi(channelIdStr)
	channelInfo := &usersModel.Channel{}
	result := database.Db.Model(channelInfo).Where("symbol = ?", channelSymbol).Where("id = ?", channelId).Where("admin_id = ?", ctx.AdminSettingId).Find(channelInfo)
	if result.Error != nil || channelInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	userInfo := &usersModel.User{}
	database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Find(userInfo)

	// 授权充值
	depositParams := &usersService.ApproveDeposit{Symbol: channelInfo.Symbol, User: userInfo.UserName, Money: params.Money, Time: time.Now().Unix()}
	err := usersService.NewUserChannel(channelInfo).ApproveDeposit(depositParams)
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}

	// 添加当前用户金额
	err = database.Db.Transaction(func(tx *gorm.DB) error {
		result = tx.Create(&walletsModel.WalletUserOrder{
			AdminId:  ctx.AdminId,
			UserId:   ctx.UserId,
			AssetsId: paymentInfo.AssetsId,
			Type:     paymentInfo.Mode,
			SourceId: paymentInfo.ID,
			OrderSn:  utils.NewRandom().OrderSn(), Money: params.Money, Voucher: params.Voucher, Status: walletsModel.WalletUserOrderStatusComplete,
		})
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}

		// 用户添加余额
		return walletsService.NewUserWallet(tx, userInfo, nil).ChangeUserBalance(walletsModel.WalletUserBillTypeChannelDeposit, paymentInfo.ID, params.Money)
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}
	return ctx.SuccessJsonOK()
}
