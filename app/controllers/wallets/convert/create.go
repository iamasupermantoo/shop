package convert

import (
	"errors"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
)

type CreateParams struct {
	AssetsId    uint    `json:"assetsId"`                       //	资产ID(出)
	ToAssetsId  uint    `json:"toAssetsId"`                     //	资产ID(入)
	Money       float64 `json:"money" validate:"required,gt=0"` //	数量
	SecurityKey string  `json:"securityKey"`                    // 安全密钥
}

// Create 资金转换申请
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	userInfo := &usersModel.User{}
	database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Find(userInfo)

	adminSettingCache := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId)
	walletsTemplate := adminSettingCache.CheckBoxToMaps("walletsTemplate")
	// 判断需要安全密钥
	if walletsTemplate["showConvertSecurityKey"] && userInfo.SecurityKey != utils.PasswordEncrypt(params.SecurityKey) {
		return ctx.ErrorJsonTranslate("incorrectSecurityKey")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		convertInfo := &walletsModel.WalletUserConvert{
			AdminId:    userInfo.AdminId,
			UserId:     userInfo.ID,
			AssetsId:   params.AssetsId,
			ToAssetsId: params.ToAssetsId,
			Money:      params.Money,
		}
		result := tx.Create(convertInfo)
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}

		var err error
		if params.AssetsId == 0 || params.ToAssetsId == 0 {
			// 余额和资产之间闪兑
			err = convertBalanceToAssets(tx, ctx.AdminSettingId, userInfo, params, convertInfo)
		} else {
			// 资产之间闪兑
			err = convertAssetsToAssets(tx, ctx.AdminSettingId, userInfo, params, convertInfo)
		}

		return err
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	return ctx.SuccessJsonOK()
}

// 余额资产转换
func convertBalanceToAssets(tx *gorm.DB, settingId uint, userInfo *usersModel.User, params *CreateParams, convertInfo *walletsModel.WalletUserConvert) error {
	walletsInstance := walletsService.NewUserWallet(tx, userInfo, nil)
	var err error
	if params.AssetsId == 0 {
		// 余额兑换资产
		toAssetsInfo := &walletsModel.WalletAssets{}
		result := database.Db.Model(toAssetsInfo).Where("id = ?", params.ToAssetsId).Where("admin_id = ?", settingId).Find(toAssetsInfo)
		if result.Error != nil || toAssetsInfo.ID == 0 {
			return errors.New("abnormalOperation")
		}
		assetsNums := params.Money / toAssetsInfo.Rate

		// 更新转换订单数量
		tx.Model(&walletsModel.WalletUserConvert{}).Where("id = ?", convertInfo.ID).Update("nums", assetsNums)

		err = walletsInstance.ChangeUserBalance(walletsModel.WalletUserBillTypeConvertWithdraw, convertInfo.ID, params.Money)
		if err != nil {
			return err
		}
		err = walletsInstance.SetAssetsInfo(toAssetsInfo).ChangeUserAssets(walletsModel.WalletUserBillTypeAssetsConvertDeposit, convertInfo.ID, assetsNums)

	} else {
		//	资产兑换余额
		assetsInfo := &walletsModel.WalletAssets{}
		result := database.Db.Model(assetsInfo).Where("id = ?", params.AssetsId).Where("admin_id = ?", settingId).Find(assetsInfo)
		if result.Error != nil || assetsInfo.ID == 0 {
			return errors.New("abnormalOperation")
		}

		balanceNums := params.Money * assetsInfo.Rate
		// 更新转换订单数量
		tx.Model(&walletsModel.WalletUserConvert{}).Where("id = ?", convertInfo.ID).Update("nums", balanceNums)

		err = walletsInstance.SetAssetsInfo(assetsInfo).ChangeUserAssets(walletsModel.WalletUserBillTypeAssetsConvertWithdraw, convertInfo.ID, params.Money)
		if err != nil {
			return err
		}
		err = walletsInstance.ChangeUserBalance(walletsModel.WalletUserBillTypeConvertDeposit, convertInfo.ID, balanceNums)
	}
	return err
}

// 资产之间转换
func convertAssetsToAssets(tx *gorm.DB, settingId uint, userInfo *usersModel.User, params *CreateParams, convertInfo *walletsModel.WalletUserConvert) error {
	walletsInstance := walletsService.NewUserWallet(tx, userInfo, nil)
	assetsInfo := &walletsModel.WalletAssets{}
	result := database.Db.Model(assetsInfo).Where("id = ?", params.AssetsId).Where("admin_id = ?", settingId).Find(assetsInfo)
	if result.Error != nil || assetsInfo.ID == 0 {
		return errors.New("abnormalOperation")
	}
	toAssetsInfo := &walletsModel.WalletAssets{}
	result = database.Db.Model(toAssetsInfo).Where("id = ?", params.ToAssetsId).Where("admin_id = ?", settingId).Find(toAssetsInfo)
	if result.Error != nil || toAssetsInfo.ID == 0 {
		return errors.New("abnormalOperation")
	}

	// 兑换数量
	assetsNums := assetsInfo.Rate * params.Money / toAssetsInfo.Rate

	// 更新转换订单数量
	tx.Model(&walletsModel.WalletUserConvert{}).Where("id = ?", convertInfo.ID).Update("nums", assetsNums)

	// 支出资产数量
	err := walletsInstance.SetAssetsInfo(assetsInfo).ChangeUserAssets(walletsModel.WalletUserBillTypeAssetsConvertWithdraw, convertInfo.ID, params.Money)
	if err != nil {
		return err
	}

	// 收入资产数量
	return walletsInstance.SetAssetsInfo(toAssetsInfo).ChangeUserAssets(walletsModel.WalletUserBillTypeAssetsConvertDeposit, convertInfo.ID, assetsNums)
}
