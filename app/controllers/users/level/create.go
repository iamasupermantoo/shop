package level

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
	"time"
)

type CreateParams struct {
	ID int `json:"id" validate:"required"` //等级ID
}

// Create 购买等级
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	// 获取当前用户信息
	userInfo := &usersModel.User{}
	result := database.Db.Model(userInfo).Where(ctx.UserId).Find(userInfo)
	if result.Error != nil {
		return ctx.ErrorJson("abnormalOperation")
	}

	//获取购买等级信息
	levelInfo := &systemsModel.Level{}
	result = database.Db.Model(levelInfo).Where(params.ID).Where("status = ?", systemsModel.LevelStatusActive).Find(levelInfo)
	if result.Error != nil || levelInfo.ID == 0 {
		return ctx.ErrorJson("abnormalOperation")
	}
	buyMoney := levelInfo.Money
	buyDays := levelInfo.Days
	if buyDays == -1 {
		buyDays = 365
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		userLevelInfo := &usersModel.UserLevel{}
		result = database.Db.Model(userLevelInfo).Where("user_id = ?", userInfo.ID).Find(userLevelInfo)
		if result.RowsAffected == 0 && userLevelInfo.ID == 0 {
			userLevelInfo = &usersModel.UserLevel{
				AdminId:   userInfo.AdminId,
				UserId:    userInfo.ID,
				Name:      levelInfo.Name,
				Symbol:    levelInfo.Symbol,
				Icon:      levelInfo.Icon,
				Money:     levelInfo.Money,
				ExpiredAt: time.Now().Add(time.Duration(buyDays) * 24 * time.Hour),
			}
			result = tx.Create(userLevelInfo)
		} else {
			// 必须大于当前标识, 可以购买成功
			if userLevelInfo.Symbol >= levelInfo.Symbol {
				return ctx.ErrorJsonTranslate("lowerLevel")
			}

			if levelInfo.Type == systemsModel.LevelTypeDifference {
				buyMoney = buyMoney - userLevelInfo.Money
			}
			result = tx.Model(&usersModel.UserLevel{}).Where("id = ?", userLevelInfo.ID).Updates(&usersModel.UserLevel{
				Name: levelInfo.Name, Icon: levelInfo.Icon, Symbol: levelInfo.Symbol, Money: levelInfo.Money, ExpiredAt: time.Now().Add(time.Duration(buyDays) * 24 * time.Hour),
			})
		}
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}

		// 资金扣款
		return walletsService.NewUserWallet(tx, userInfo, nil).ChangeUserBalance(walletsModel.WalletUserBillTypeBuyLevel, userLevelInfo.ID, buyMoney)
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	return ctx.SuccessJsonOK()
}
