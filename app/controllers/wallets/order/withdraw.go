package order

import (
	"github.com/goccy/go-json"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/websocket"
	"gofiber/utils"
	"gorm.io/gorm"
	"time"
)

type WithdrawParams struct {
	AccountId   uint    `json:"accountId" validate:"required"`  // 提现账户ID
	Money       float64 `json:"money" validate:"required,gt=0"` // 金额
	SecurityKey string  `json:"securityKey"`                    // 安全密钥
}

type withdrawValue struct {
	Days uint    `json:"days"` //	天数
	Nums uint    `json:"nums"` //	次数
	Fee  float64 `json:"fee"`  //	手续费
}

// Withdraw 钱包订单提现
func Withdraw(ctx *context.CustomCtx, params *WithdrawParams) error {
	userInfo := &usersModel.User{}
	database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Find(userInfo)

	// 获取用户账户和体现方式
	accountInfo := &walletUserAccount{}
	result := database.Db.Model(&walletsModel.WalletUserAccount{}).Where("id = ?", params.AccountId).
		Preload("PaymentInfo").
		Where("user_id = ?", userInfo.ID).Where("status = ?", walletsModel.WalletUserAccountStatusActive).Find(accountInfo)
	if result.Error != nil || (accountInfo.PaymentInfo.Mode != walletsModel.WalletPaymentModeWithdraw && accountInfo.PaymentInfo.Mode != walletsModel.WalletPaymentModeAssetsWithdraw) {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}
	adminSettingService := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId)
	walletsTemplate := adminSettingService.CheckBoxToMaps("walletsTemplate")
	freezeTemplate := adminSettingService.CheckBoxToMaps("freezeTemplate")

	// 如果用户冻结状态
	if userInfo.Status == usersModel.UserStatusDisable && freezeTemplate["closeWithdraw"] {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	// 判断需要安全密钥
	if walletsTemplate["showWithdrawSecurityKey"] && userInfo.SecurityKey != utils.PasswordEncrypt(params.SecurityKey) {
		return ctx.ErrorJsonTranslate("incorrectSecurityKey")
	}

	moneyRate := accountInfo.PaymentInfo.Rate * params.Money
	rangeMoney := adminSettingService.GetRangeMoney("walletWithdrawAmountBetween")
	if moneyRate < rangeMoney.Min || moneyRate > rangeMoney.Max {
		return ctx.ErrorJsonTranslate("rangeMismatch", rangeMoney.Min, rangeMoney.Max)
	}

	// 用户有未完成的订单，返回错误
	if result = database.Db.Where("user_id = ?", ctx.UserId).
		Where("status = ?", walletsModel.WalletUserOrderStatusActive).
		Where("type = ?", accountInfo.PaymentInfo.Mode).
		Find(&walletsModel.WalletUserOrder{}); result.RowsAffected != 0 {
		return ctx.ErrorJsonTranslate("limitExceeded")
	}

	// 根据管理提现设置的value判断能否继续提现
	walletWithdrawSettingValue := adminSettingService.GetRedisAdminSettingField("walletWithdrawSetting")
	var withdrawValueStruct withdrawValue
	err := json.Unmarshal([]byte(walletWithdrawSettingValue), &withdrawValueStruct)
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}

	//	根据提现设置中的天数和提现次数判断能否继续提现
	currentTime := time.Now()
	var num int64
	if result := database.Db.Model(&walletsModel.WalletUserOrder{}).
		Where("admin_id = ?", ctx.AdminId).
		Where("created_at between ? and ?", currentTime.AddDate(0, 0, -int(withdrawValueStruct.Days)), currentTime).
		Where("type = ? or type = ?", walletsModel.WalletUserOrderTypeWithdraw, walletsModel.WalletUserOrderTypeAssetsWithdraw).
		Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
		Count(&num); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}
	if uint(num) >= withdrawValueStruct.Nums {
		return ctx.ErrorJsonTranslate("limitExceeded")
	}

	// 事务操作
	err = database.Db.Transaction(func(tx *gorm.DB) error {
		orderType := walletsModel.WalletPaymentModeWithdraw
		if accountInfo.PaymentInfo.Mode == walletsModel.WalletPaymentModeAssetsWithdraw {
			orderType = walletsModel.WalletPaymentModeAssetsWithdraw
		}
		orderInfo := &walletsModel.WalletUserOrder{
			AdminId:  userInfo.AdminId,
			UserId:   userInfo.ID,
			AssetsId: accountInfo.PaymentInfo.AssetsId,
			SourceId: accountInfo.ID,
			OrderSn:  utils.NewRandom().OrderSn(),
			Money:    params.Money,
			Fee:      withdrawValueStruct.Fee,
			Type:     orderType,
		}
		result = tx.Create(orderInfo)
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}
		walletsInstance := walletsService.NewUserWallet(tx, userInfo, nil)

		var err error
		switch accountInfo.PaymentInfo.Mode {
		// 余额提现
		case walletsModel.WalletPaymentModeWithdraw:
			err = walletsInstance.ChangeUserBalance(walletsModel.WalletUserBillTypeWithdraw, orderInfo.ID, params.Money+withdrawValueStruct.Fee)
		// 资产提现
		case walletsModel.WalletPaymentModeAssetsWithdraw:
			assetsInfo := &walletsModel.WalletAssets{}
			result = database.Db.Model(assetsInfo).Where("id = ?", accountInfo.PaymentInfo.AssetsId).Find(assetsInfo)
			if result.Error != nil || assetsInfo.ID == 0 {
				return ctx.ErrorJsonTranslate("abnormalOperation")
			}
			err = walletsInstance.SetAssetsInfo(assetsInfo).ChangeUserAssets(walletsModel.WalletUserBillTypeAssetsWithdraw, orderInfo.ID, params.Money+withdrawValueStruct.Fee)
		}
		return err
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	// 管理通知
	websocket.NewAdminNotify(ctx.Rds, ctx.AdminId).Deposit(*userInfo, accountInfo.PaymentInfo, params.Money)
	return ctx.SuccessJsonOK()
}

type walletUserAccount struct {
	walletsModel.WalletUserAccount
	PaymentInfo walletsModel.WalletPayment ` json:"paymentInfo" gorm:"foreignKey:PaymentId"`
}
