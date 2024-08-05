package user

import (
	"github.com/dchest/captcha"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/commonService"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
	"strings"
	"time"
)

type RegisterParams struct {
	UserName    string `json:"userName" validate:"required,alphanum"` //	用户名
	Password    string `json:"password" validate:"required"`          //	密码
	CaptchaId   string `json:"captchaId"`                             //	验证码ID
	CaptchaVal  string `json:"captchaVal"`                            //	验证码值
	Email       string `json:"email"`                                 //	邮箱
	Telephone   string `json:"telephone"`                             //	手机号码
	SecurityKey string `json:"securityKey"`                           //	安全密钥
	Code        string `json:"code"`                                  //	用户邀请码
}

// Register 用户注册
func Register(ctx *context.CustomCtx, params *RegisterParams) error {
	adminSettingCache := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId)
	registerTemplate := adminSettingCache.CheckBoxToMaps("registerTemplate")

	// 如果需要验证码
	if registerTemplate["showVerify"] {
		if !captcha.VerifyString(params.CaptchaId, params.CaptchaVal) {
			return ctx.ErrorJsonTranslate("codeError")
		}
	}

	// 是否需要邮箱
	if registerTemplate["showEmail"] && params.Email == "" {
		return ctx.ErrorJsonTranslateMultiple("email", "notBeEmpty")
	}

	// 手机号码不能为空
	if registerTemplate["showTelephone"] && params.Telephone == "" {
		return ctx.ErrorJsonTranslateMultiple("telephone", "notBeEmpty")
	}

	// 安全密钥不能为空
	if registerTemplate["showSecurityKey"] && params.SecurityKey == "" {
		return ctx.ErrorJsonTranslateMultiple("secretKey", "notBeEmpty")
	}

	// 显示邀请码
	if registerTemplate["showInvite"] && params.Code == "" {
		return ctx.ErrorJsonTranslateMultiple("inviteCode", "notBeEmpty")
	}

	adminId := ctx.AdminId
	var userParentId uint
	if params.Code != "" {
		inviteInfo := &usersModel.Invite{}
		database.Db.Model(inviteInfo).Where("code = ?", params.Code).Where("status = ?", usersModel.InviteStatusActive).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(inviteInfo)

		// 如果当前邀请码不属于同一个分支, 那么验证码不正确
		if inviteInfo.ID == 0 {
			return ctx.ErrorJsonTranslate("codeError")
		}
		adminId = inviteInfo.AdminId
		userParentId = inviteInfo.UserId
	}

	// 监测当前用户名是否存在
	createdUserInfo := &usersModel.User{}
	result := database.Db.Model(createdUserInfo).Where("user_name = ?", params.UserName).Find(createdUserInfo)
	if result.RowsAffected > 0 {
		return ctx.ErrorJsonTranslate("accountAlreadyExists")
	}
	//	插入用户数据
	if params.SecurityKey != "" {
		params.SecurityKey = utils.PasswordEncrypt(params.SecurityKey)
	}

	createdUserInfo = &usersModel.User{
		ParentId:    userParentId,
		AdminId:     adminId,
		UserName:    strings.TrimSpace(params.UserName),
		Password:    utils.PasswordEncrypt(strings.TrimSpace(params.Password)),
		SecurityKey: params.SecurityKey,
		Telephone:   params.Telephone,
		Email:       params.Email,
		Birthday:    time.Now().Add(-24 * 365 * 24 * time.Hour),
	}
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		result = tx.Create(createdUserInfo)
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation")
		}

		// 邀请奖励 - 分销奖励
		registerAward := adminSettingCache.GetRegisterAward()

		// 自身奖励
		userWallet := walletsService.NewUserWallet(tx, createdUserInfo, nil)
		if registerAward.Register > 0 {
			err := userWallet.ChangeUserBalance(walletsModel.WalletUserBillTypeRegisterAward, 0, registerAward.Register)
			if err != nil {
				return err
			}
		}
		// 邀请者奖励
		if registerAward.Share > 0 && userParentId > 0 {
			userParentInfo := &usersModel.User{}
			database.Db.Model(userParentInfo).Where("id = ?", userParentId).Find(userParentInfo)
			err := userWallet.ChangeUserBalance(walletsModel.WalletUserBillTypeShareAward, createdUserInfo.ID, registerAward.Share)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	userInfo := &homeUserInfo{}
	database.Db.Model(userInfo).Where("id = ?", createdUserInfo.ID).Preload("AuthInfo").Preload("LevelInfo").Find(userInfo)

	return ctx.SuccessJson(authData{
		UserInfo: userInfo,
		Token:    commonService.NewServiceToken(ctx.Rds).GenerateHomeToken(userInfo.AdminId, userInfo.ID),
	})
}
