package settled

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
	"time"
)

// StatusParams	店铺入驻审核
type StatusParams struct {
	ID       uint   `json:"id" validate:"required"` // ID
	OpStatus int    `json:"opStatus"`               // 状态 -1拒绝 20通过
	Data     string `json:"data"`                   // 拒绝理由
}

// Status 店铺入驻审核
func Status(ctx *context.CustomCtx, params *StatusParams) error {
	// 判断入驻信息是否存在
	settledInfo := &shopsModel.StoreSettled{}
	result := database.Db.Where("id = ?", params.ID).Where("admin_id in ?", ctx.GetAdminChildIds()).Find(&settledInfo)
	if result.Error != nil || settledInfo.ID == 0 {
		return ctx.ErrorJson("找不到入驻信息")
	}

	switch params.OpStatus {
	case shopsModel.StoreSettledStatusPass:
		err := database.Db.Transaction(func(tx *gorm.DB) error {
			result = tx.Model(&shopsModel.StoreSettled{}).Where("id = ?", settledInfo.ID).Update("status", shopsModel.StoreSettledStatusPass)
			if result.Error != nil {
				return ctx.ErrorJson("更新入驻信息失败" + "[" + result.Error.Error() + "]")
			}

			// 更新当前用户信息
			userInfo := &usersModel.User{}
			result = database.Db.Model(userInfo).Where("id = ?", settledInfo.UserId).Find(userInfo)
			if result.Error != nil || userInfo.ID == 0 {
				return ctx.ErrorJson("找不到用户信息" + "[" + result.Error.Error() + "]")
			}

			// 开启当前店铺
			result = tx.Create(&shopsModel.Store{
				AdminId: settledInfo.AdminId, UserId: settledInfo.UserId, Logo: settledInfo.Logo, Name: settledInfo.Name,
				Contact: settledInfo.Contact, Address: settledInfo.Address, Status: shopsModel.StoreStatusActivate,
			})
			if result.Error != nil {
				return ctx.ErrorJson("开启店铺失败" + "[" + result.Error.Error() + "]")
			}

			// 更新用户信息
			result = tx.Model(&usersModel.User{}).Where("id = ?", settledInfo.UserId).Updates(&usersModel.User{
				CountryId: settledInfo.CountryId, NickName: settledInfo.RealName, Telephone: settledInfo.Contact,
				Email: settledInfo.Email,
			})
			if result.Error != nil {
				return ctx.ErrorJson("更新用户信息失败" + "[" + result.Error.Error() + "]")
			}

			// 更新实名信息通过
			userAuthInfo := &usersModel.UserAuth{}
			result = database.Db.Model(userAuthInfo).Where("user_id = ?", settledInfo.UserId).Find(userAuthInfo)
			if result.Error == nil && userAuthInfo.ID == 0 {
				result = tx.Create(&usersModel.UserAuth{
					AdminId: settledInfo.AdminId, UserId: settledInfo.UserId, RealName: settledInfo.RealName, Number: settledInfo.Number,
					Photo1: settledInfo.Photo1, Photo2: settledInfo.Photo2, Photo3: settledInfo.Photo3, Address: settledInfo.Address,
					Type: settledInfo.Type, Status: usersModel.UserAuthStatusComplete,
				})
				if result.Error != nil {
					return ctx.ErrorJson("同步实名认证失败" + "[" + result.Error.Error() + "]")
				}
			}

			// 购买第一个会员
			currentLevelInfo := &systemsModel.Level{}
			tx.Model(currentLevelInfo).Where("admin_id = ?", ctx.AdminSettingId).Where("status = ?", systemsModel.LevelStatusActive).
				Order("symbol ASC").Limit(1).Find(currentLevelInfo)
			if currentLevelInfo.ID > 0 {
				expireTime := time.Now()
				if currentLevelInfo.Days == -1 {
					expireTime = expireTime.Add(365 * 24 * time.Hour)
				} else {
					expireTime = expireTime.Add(time.Duration(currentLevelInfo.Days) * 24 * time.Hour)
				}

				result = tx.Create(&usersModel.UserLevel{
					AdminId: userInfo.AdminId, UserId: userInfo.ID, Name: currentLevelInfo.Name, Icon: currentLevelInfo.Icon,
					Symbol: currentLevelInfo.Symbol, Money: currentLevelInfo.Money, ExpiredAt: expireTime, Increase: currentLevelInfo.Increase,
				})
			}

			return nil
		})
		if err != nil {
			return ctx.ErrorJson(err.Error())
		}
	default:
		if params.Data == "" {
			return ctx.ErrorJson("请输入拒绝理由")
		}

		result = database.Db.Model(&shopsModel.StoreSettled{}).Where("id = ?", settledInfo.ID).Updates(&shopsModel.StoreSettled{
			Status: shopsModel.StoreSettledStatusRefuse, Data: params.Data,
		})
		if result.Error != nil {
			return ctx.ErrorJson("更新拒绝失败" + "[" + result.Error.Error() + "]")
		}
	}
	return ctx.SuccessJsonOK()
}
