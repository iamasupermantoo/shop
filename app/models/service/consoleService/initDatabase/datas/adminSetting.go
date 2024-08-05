package datas

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/views"
	"gofiber/utils"
)

// InitAdminSetting 初始化管理配置
func InitAdminSetting() (adminSetting []*adminsModel.AdminSetting) {
	// 基础配置
	adminSetting = append(adminSetting, []*adminsModel.AdminSetting{
		{Name: "站点Logo", Field: "siteLogo", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupBasic,
			Type:  views.InputTypeImage,
			Value: "/logo.png",
		},

		{Name: "站点名称", Field: "siteName", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupBasic,
			Type:  views.InputTypeText,
			Value: "Tiktok",
		},

		{Name: "轮播图", Field: "siteBanners", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupBasic,
			Type: views.InputTypeImages,
			Value: utils.JsonMarshal([]string{
				"/assets/banner/banner1.png",
				"/assets/banner/banner2.jpeg",
				"/assets/banner/banner3.jpeg",
				"/assets/banner/banner4.jpeg",
				"/assets/banner/banner5.jpeg",
				"/assets/banner/banner6.jpeg",
			})},

		{Name: "提示音设置", Field: "audioSetting", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupBasic,
			Type: views.InputTypeCheckbox,
			Value: utils.JsonMarshal([]*views.InputCheckboxOptions{
				{Label: "余额充值", Value: false, Field: "balanceDeposit"},
				{Label: "资产充值", Value: false, Field: "assetsDeposit"},
				{Label: "余额提现", Value: false, Field: "balanceWithdraw"},
				{Label: "资产提现", Value: false, Field: "assetsWithdraw"},
				{Label: "下单提示", Value: false, Field: "createOrder"},
			}),
			Data: utils.JsonMarshal([]*views.InputOptions{
				{Label: "余额充值", Value: false}, {Label: "资产充值", Value: false},
				{Label: "余额提现", Value: false}, {Label: "资产提现", Value: false},
				{Label: "下单提示", Value: false},
			}),
		},

		{Name: "站点信息", Field: "siteInfo", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupBasic,
			Type: views.InputTypeJson,
			Value: utils.JsonMarshal(&adminsModel.AdminSettingSiteInfo{
				Introduce: "siteIntroduce",
				Notice:    "siteNotice",
			}),
			Data: utils.JsonMarshal([][]*views.InputAttrsViews{{
				{Label: "站点信息", Field: "Introduce", Type: views.InputTypeTranslate},
				{Label: "站点公告", Field: "Notice", Type: views.InputTypeTranslate},
			}})},

		{Name: "软件下载", Field: "download", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupBasic,
			Type: views.InputTypeJson,
			Value: utils.JsonMarshal(&adminsModel.AdminSettingDownload{
				Android: "",
				Ios:     "",
			}),
			Data: utils.JsonMarshal([][]*views.InputAttrsViews{{
				{Label: "安卓包地址", Field: "android", Type: views.InputTypeFile},
				{Label: "苹果包地址", Field: "ios", Type: views.InputTypeFile},
			}}),
		},
	}...)

	// 钱包配置
	adminSetting = append(adminSetting, []*adminsModel.AdminSetting{
		{Name: "绑定账户数量", Field: "walletAccountNums", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupWallet,
			Type:  views.InputTypeNumber,
			Value: "5",
		},

		{Name: "货币汇率", Field: "amountRate", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupWallet,
			Type:  views.InputTypeNumber,
			Value: "0.98",
		},

		{Name: "充值范围", Field: "walletDepositAmountBetween", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupWallet,
			Type: views.InputTypeJson,
			Value: utils.JsonMarshal(&adminsModel.AdminSettingRange{
				Max: 99999999,
				Min: 100,
			}),
			Data: utils.JsonMarshal([][]*views.InputAttrsViews{{
				{Label: "最小金额", Field: "min", Type: views.InputTypeNumber},
				{Label: "最大金额", Field: "max", Type: views.InputTypeNumber},
			}}),
		},

		{Name: "充值提示", Field: "walletDepositDesc", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupWallet,
			Type:  views.InputTypeTranslate,
			Value: "depositDesc",
		},

		{Name: "提现范围", Field: "walletWithdrawAmountBetween", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupWallet,
			Type: views.InputTypeJson,
			Value: utils.JsonMarshal(&adminsModel.AdminSettingRange{
				Max: 99999999,
				Min: 100,
			}),
			Data: utils.JsonMarshal([][]*views.InputAttrsViews{{
				{Label: "最小金额", Field: "min", Type: views.InputTypeNumber},
				{Label: "最大金额", Field: "max", Type: views.InputTypeNumber},
			}}),
		},

		{Name: "提现提示", Field: "walletWithdrawDesc", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupWallet,
			Type:  views.InputTypeTranslate,
			Value: "withdrawDesc",
		},

		{Name: "提现配置", Field: "walletWithdrawSetting", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupWallet,
			Type: views.InputTypeJson,
			Value: utils.JsonMarshal(&adminsModel.AdminSettingWithdraw{
				Days: 7,
				Nums: 7,
				Fee:  0,
			}),
			Data: utils.JsonMarshal([][]*views.InputAttrsViews{{
				{Label: "天数", Field: "days", Type: views.InputTypeNumber},
				{Label: "次数", Field: "nums", Type: views.InputTypeNumber},
				{Label: "手续费", Field: "fee", Type: views.InputTypeNumber},
			}}),
		},

		{Name: "转移手续费", Field: "transferFee", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupWallet,
			Type:  views.InputTypeNumber,
			Value: "0",
		},

		{Name: "推广奖励", Field: "registerAward", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupWallet,
			Type: views.InputTypeJson,
			Value: utils.JsonMarshal(&adminsModel.AdminSettingRegisterAward{
				Register: 0,
				Share:    0,
			}),
			Data: utils.JsonMarshal([][]*views.InputAttrsViews{{
				{Label: "注册奖励", Field: "register", Type: views.InputTypeNumber},
				{Label: "邀请奖励", Field: "share", Type: views.InputTypeNumber},
			}}),
		},

		{Name: "分销配置", Field: "earningsSetting", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupWallet,
			Type: views.InputTypeChildren,
			Value: utils.JsonMarshal([]*adminsModel.AdminSettingEarningsSetting{
				{Options: []*views.InputCheckboxOptions{
					{Label: "购买产品", Value: false, Field: "buyProduct"},
					{Label: "产品收益", Value: false, Field: "productEarnings"},
				}, Rate: 50},
				{Options: []*views.InputCheckboxOptions{
					{Label: "购买产品", Value: false, Field: "buyProduct"},
					{Label: "产品收益", Value: false, Field: "productEarnings"},
				}, Rate: 30},
				{Options: []*views.InputCheckboxOptions{
					{Label: "购买产品", Value: false, Field: "buyProduct"},
					{Label: "产品收益", Value: false, Field: "productEarnings"},
				}, Rate: 10},
			}),
			Data: utils.JsonMarshal([][]*views.InputAttrsViews{{
				{Label: "功能选项", Field: "options", Type: views.InputTypeCheckbox, Data: []*views.InputOptions{{Label: "购买产品", Value: false}, {Label: "产品收益", Value: false}}},
				{Label: "收益率(%)", Field: "rate", Type: views.InputTypeNumber},
			}}),
		},
	}...)

	// 项目基础模版
	adminSetting = append(adminSetting, []*adminsModel.AdminSetting{
		{Name: "基础模版", Field: "basicTemplate", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupTemplate,
			Type: views.InputTypeCheckbox,
			Value: utils.JsonMarshal([]*views.InputCheckboxOptions{
				{Label: "显示实名图标", Value: true, Field: "showAuth"},
				{Label: "显示等级图标", Value: true, Field: "showLevel"},
				{Label: "显示信用分", Value: true, Field: "showScore"},
			}),
			Data: utils.JsonMarshal([]*views.InputOptions{
				{Label: "显示实名图标", Value: false},
				{Label: "显示等级图标", Value: false},
				{Label: "显示信用分", Value: false},
			}),
		},

		{Name: "实名模版", Field: "authTemplate", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupTemplate,
			Type: views.InputTypeCheckbox,
			Value: utils.JsonMarshal([]*views.InputCheckboxOptions{
				{Label: "显示证件名称", Value: true, Field: "showRealName"},
				{Label: "显示证件卡号", Value: true, Field: "showNumber"},
				{Label: "显示证件照2", Value: true, Field: "showPhoto2"},
				{Label: "显示证件照3", Value: false, Field: "showPhoto3"},
				{Label: "显示详细地址", Value: true, Field: "showAddress"},
			}),
			Data: utils.JsonMarshal([]*views.InputOptions{
				{Label: "显示证件名称", Value: true},
				{Label: "显示证件卡号", Value: true},
				{Label: "显示证件照2", Value: true},
				{Label: "显示证件照3", Value: true},
				{Label: "显示详细地址", Value: true},
			}),
		},

		{Name: "注册模版", Field: "registerTemplate", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupTemplate,
			Type: views.InputTypeCheckbox,
			Value: utils.JsonMarshal([]*views.InputCheckboxOptions{
				{Label: "显示电子邮箱", Value: false, Field: "showEmail"},
				{Label: "显示手机号码", Value: false, Field: "showTelephone"},
				{Label: "显示安全密码", Value: false, Field: "showSecurityKey"},
				{Label: "显示确认安全密码", Value: false, Field: "showCmfSecurityKey"},
				{Label: "显示确认密码", Value: false, Field: "showCmfPass"},
				{Label: "显示验证码", Value: true, Field: "showVerify"},
				{Label: "显示邀请码", Value: false, Field: "showInvite"},
				{Label: "显示切换语言", Value: true, Field: "showLang"},
				{Label: "显示在线客服", Value: true, Field: "showOnline"},
			}),
			Data: utils.JsonMarshal([]*views.InputOptions{
				{Label: "显示电子邮箱", Value: false},
				{Label: "显示手机号码", Value: false},
				{Label: "显示安全密码", Value: false},
				{Label: "显示确认安全密码", Value: false},
				{Label: "显示确认密码", Value: false},
				{Label: "显示验证码", Value: false},
				{Label: "显示邀请码", Value: false},
				{Label: "显示切换语言", Value: false},
				{Label: "显示在线客服", Value: false},
			}),
		},

		{Name: "登录模版", Field: "loginTemplate", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupTemplate,
			Type: views.InputTypeCheckbox,
			Value: utils.JsonMarshal([]*views.InputCheckboxOptions{
				{Label: "显示验证码", Value: true, Field: "showVerify"},
				{Label: "显示切换语言", Value: true, Field: "showLang"},
				{Label: "显示在线客服", Value: true, Field: "showOnline"},
				{Label: "显示注册按钮", Value: true, Field: "showRegister"},
			}),
			Data: utils.JsonMarshal([]*views.InputOptions{
				{Label: "显示验证码", Value: false},
				{Label: "显示切换语言", Value: false},
				{Label: "显示在线客服", Value: false},
				{Label: "显示注册按钮", Value: false},
			}),
		},
		{Name: "钱包模版", Field: "walletsTemplate", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupTemplate,
			Type: views.InputTypeCheckbox,
			Value: utils.JsonMarshal([]*views.InputCheckboxOptions{
				{Label: "显示账户安全密码", Value: true, Field: "showAccountSecurityKey"},
				{Label: "显示账户更新", Value: true, Field: "showAccountUpdate"},
				{Label: "显示账户删除", Value: true, Field: "showAccountDelete"},
				{Label: "显示账户卡号", Value: true, Field: "showAccountNumber"},
				{Label: "显示提现安全密钥", Value: true, Field: "showWithdrawSecurityKey"},
				{Label: "显示闪兑安全密钥", Value: true, Field: "showConvertSecurityKey"},
				{Label: "显示转移安全密钥", Value: true, Field: "showTransferSecurityKey"},
				{Label: "显示购买安全密钥", Value: true, Field: "showOrderSecurityKey"},
			}),
			Data: utils.JsonMarshal([]*views.InputOptions{
				{Label: "显示账户安全密码", Value: false},
				{Label: "显示账户更新", Value: false},
				{Label: "显示账户删除", Value: false},
				{Label: "显示账户卡号", Value: false},
				{Label: "显示提现安全密钥", Value: false},
				{Label: "显示闪兑安全密钥", Value: false},
				{Label: "显示转移安全密钥", Value: false},
				{Label: "显示购买安全密钥", Value: false},
			}),
		},
		{Name: "首页模版", Field: "homeTemplate", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupTemplate,
			Type: views.InputTypeCheckbox,
			Value: utils.JsonMarshal([]*views.InputCheckboxOptions{
				{Label: "显示切换语言", Value: true, Field: "showLang"},
				{Label: "显示在线客服", Value: true, Field: "showOnline"},
			}),
			Data: utils.JsonMarshal([]*views.InputOptions{
				{Label: "显示切换语言", Value: true},
				{Label: "显示在线客服", Value: true},
			}),
		},
		{Name: "用户冻结", Field: "freezeTemplate", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupTemplate,
			Type: views.InputTypeCheckbox,
			Value: utils.JsonMarshal([]*views.InputCheckboxOptions{
				{Label: "禁止登录", Value: true, Field: "closeLogin"},
				{Label: "禁止提现", Value: true, Field: "closeWithdraw"},
				{Label: "禁止下单", Value: true, Field: "closeOrder"},
			}),
			Data: utils.JsonMarshal([]*views.InputOptions{
				{Label: "禁止登录", Value: false},
				{Label: "禁止提现", Value: false},
				{Label: "禁止下单", Value: false},
			}),
		},
		{Name: "入驻模版", Field: "settledTemplate", AdminId: adminsModel.SuperAdminId, GroupId: adminsModel.AdminSettingGroupTemplate,
			Type: views.InputTypeCheckbox,
			Value: utils.JsonMarshal([]*views.InputCheckboxOptions{
				{Label: "显示国家", Value: false, Field: "showCountry"},
				{Label: "显示证件卡号", Value: true, Field: "showNumber"},
				{Label: "显示证件名称", Value: false, Field: "showRealName"},
				{Label: "显示证件照2", Value: true, Field: "showPhoto2"},
				{Label: "显示证件照3", Value: false, Field: "showPhoto3"},
				{Label: "显示邮箱", Value: false, Field: "showEmail"},
				{Label: "显示手机号码", Value: false, Field: "showContact"},
			}),
			Data: utils.JsonMarshal([]*views.InputOptions{
				{Label: "显示证件卡号", Value: true},
				{Label: "显示证件名称", Value: true},
				{Label: "显示证件照2", Value: true},
				{Label: "显示证件照3", Value: true},
				{Label: "显示邮箱", Value: true},
				{Label: "显示手机号码", Value: true},
			})},
	}...)

	return adminSetting
}
