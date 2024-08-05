package transfer

import (
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
	AssetsId    uint    `json:"assetsId"`                        // 资产ID
	UserName    string  `json:"userName" validate:"required"`    // 用户账户
	Money       float64 ` json:"money" validate:"required,gt=0"` // 转移金额
	SecurityKey string  `json:"securityKey"`                     // 安全密钥
}

// Create 资金转移申请
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	senderInfo := &usersModel.User{}
	database.Db.Model(senderInfo).Where("id = ?", ctx.UserId).Find(senderInfo)

	receiverInfo := &usersModel.User{}
	result := database.Db.Model(receiverInfo).Where("user_name = ?", params.UserName).Where("id <> ?", ctx.UserId).Find(receiverInfo)
	if result.Error != nil || receiverInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	adminSettingCache := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId)
	walletsTemplate := adminSettingCache.CheckBoxToMaps("walletsTemplate")

	// 判断需要安全密钥
	if walletsTemplate["showTransferSecurityKey"] && senderInfo.SecurityKey != utils.PasswordEncrypt(params.SecurityKey) {
		return ctx.ErrorJsonTranslate("incorrectSecurityKey")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		transferType := walletsModel.WalletUserTransferTypeBalance
		if params.AssetsId > 0 {
			transferType = walletsModel.WalletUserTransferTypeAssets
		}

		transferInfo := &walletsModel.WalletUserTransfer{
			AdminId:    ctx.AdminId,
			SenderId:   ctx.UserId,
			ReceiverId: receiverInfo.ID,
			Type:       transferType,
			AssetsId:   params.AssetsId,
			Money:      params.Money,
		}

		result = tx.Create(transferInfo)
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}

		var err error
		walletsInstance := walletsService.NewUserWallet(tx, senderInfo, nil)

		if params.AssetsId > 0 {

			// 资产转移
			assetsInfo := &walletsModel.WalletAssets{}
			result = database.Db.Model(assetsInfo).Where("id = ?", params.AssetsId).Where("admin_id = ?", ctx.AdminSettingId).Find(assetsInfo)
			if result.Error != nil || assetsInfo.ID == 0 {
				return ctx.ErrorJsonTranslate("abnormalOperation")
			}

			// 支出
			err = walletsInstance.SetAssetsInfo(assetsInfo).ChangeUserAssets(walletsModel.WalletUserBillTypeAssetsTransferWithdraw, transferInfo.ID, params.Money)
			if err != nil {
				return err
			}

			// 收入
			err = walletsInstance.SetUserInfo(receiverInfo).ChangeUserAssets(walletsModel.WalletUserBillTypeAssetsTransferDeposit, transferInfo.ID, params.Money)
		} else {
			//	余额支出
			err = walletsInstance.SetUserInfo(senderInfo).ChangeUserBalance(walletsModel.WalletUserBillTypeTransferWithdraw, transferInfo.ID, params.Money)
			if err != nil {
				return err
			}

			//  余额收入
			err = walletsInstance.SetUserInfo(receiverInfo).ChangeUserBalance(walletsModel.WalletUserBillTypeTransferDeposit, transferInfo.ID, params.Money)
		}
		return err
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	return ctx.SuccessJsonOK()
}
