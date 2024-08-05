package walletsModel

import (
	"gofiber/app/models/model/types"
	"gofiber/app/module/views"
)

const (
	// WalletUserBillTypeDeposit 余额充值
	WalletUserBillTypeDeposit = 1

	// WalletUserBillTypeChannelDeposit 渠道充值
	WalletUserBillTypeChannelDeposit = 2

	// WalletUserBillTypeSystemDeposit 余额系统充值
	WalletUserBillTypeSystemDeposit = 3

	// WalletUserBillTypeWithdraw 余额提现
	WalletUserBillTypeWithdraw = -11

	// WalletUserBillTypeChannelWithdraw 渠道提现
	WalletUserBillTypeChannelWithdraw = -12

	// WalletUserBillTypeSystemWithdraw 余额系统扣款
	WalletUserBillTypeSystemWithdraw = -13

	// WalletUserBillTypeWithdrawRefuse 余额提现拒绝
	WalletUserBillTypeWithdrawRefuse = 15

	// WalletUserBillTypeBuyProduct 购买产品
	WalletUserBillTypeBuyProduct = -21

	// WalletUserBillTypePurchaseProduct 商家采购产品
	WalletUserBillTypePurchaseProduct = -22

	// WalletUserBillTypeBuyLevel 购买等级
	WalletUserBillTypeBuyLevel = -23

	// WalletUserBillTypeRefundProduct 产品退款
	WalletUserBillTypeRefundProduct = 31

	// WalletUserBillTypeEarnings 产品收益
	WalletUserBillTypeEarnings = 51

	// WalletUserBillTypeAward 奖励
	WalletUserBillTypeAward = 61

	// WalletUserBillTypeRegisterAward 注册奖励
	WalletUserBillTypeRegisterAward = 66

	// WalletUserBillTypeShareAward 邀请奖励
	WalletUserBillTypeShareAward = 67

	// WalletUserBillTypeTeamEarnings 团队收益
	WalletUserBillTypeTeamEarnings = 71

	// WalletUserBillTypeConvertWithdraw 余额转换提现
	WalletUserBillTypeConvertWithdraw = -81

	// WalletUserBillTypeConvertDeposit 余额转换充值
	WalletUserBillTypeConvertDeposit = 82

	// WalletUserBillTypeTransferWithdraw 余额转移提现
	WalletUserBillTypeTransferWithdraw = -91

	// WalletUserBillTypeTransferDeposit 余额转移充值
	WalletUserBillTypeTransferDeposit = 92

	// WalletUserBillTypeAssetsDeposit 资产充值
	WalletUserBillTypeAssetsDeposit = 101

	// WalletUserBillTypeSystemAssetsDeposit 资产系统充值
	WalletUserBillTypeSystemAssetsDeposit = 103

	// WalletUserBillTypeAssetsWithdraw 资产提现
	WalletUserBillTypeAssetsWithdraw = -111

	// WalletUserBillTypeSystemAssetsWithdraw 资产系统扣款
	WalletUserBillTypeSystemAssetsWithdraw = -113

	// WalletUserBillTypeAssetsWithdrawRefuse 资产提现拒绝
	WalletUserBillTypeAssetsWithdrawRefuse = 115

	// WalletUserBillTypeAssetsBuyProduct 资产购买产品
	WalletUserBillTypeAssetsBuyProduct = -121

	// WalletUserBillTypeAssetsRefundProduct 资产退款产品
	WalletUserBillTypeAssetsRefundProduct = 131

	// WalletUserBillTypeAssetsEarnings 资产产品收益
	WalletUserBillTypeAssetsEarnings = 151

	// WalletUserBillTypeAssetsAward 资产奖励
	WalletUserBillTypeAssetsAward = 161

	// WalletUserBillTypeTeamAssetsEarnings 团队资产收益
	WalletUserBillTypeTeamAssetsEarnings = 171

	// WalletUserBillTypeAssetsConvertWithdraw 资产转换提现
	WalletUserBillTypeAssetsConvertWithdraw = -201

	// WalletUserBillTypeAssetsConvertDeposit 资产转换充值
	WalletUserBillTypeAssetsConvertDeposit = 202

	// WalletUserBillTypeAssetsTransferWithdraw 资产转移提现
	WalletUserBillTypeAssetsTransferWithdraw = -221

	// WalletUserBillTypeAssetsTransferDeposit 资产转移充值
	WalletUserBillTypeAssetsTransferDeposit = 222
)

// WalletUserBillTypeList 类型分组
var WalletUserBillTypeList = []*views.InputCheckboxOptions{
	{Label: "余额充值", Field: "WalletUserBillTypeDeposit", Value: WalletUserBillTypeDeposit},
	{Label: "余额系统充值", Field: "WalletUserBillTypeSystemDeposit", Value: WalletUserBillTypeSystemDeposit},
	{Label: "余额提现拒绝", Field: "WalletUserBillTypeWithdraw", Value: WalletUserBillTypeWithdraw},
	{Label: "余额系统扣款", Field: "WalletUserBillTypeSystemWithdraw", Value: WalletUserBillTypeSystemWithdraw},
	{Label: "余额提现拒绝", Field: "WalletUserBillTypeWithdrawRefuse", Value: WalletUserBillTypeWithdrawRefuse},
	{Label: "购买产品", Field: "WalletUserBillTypeBuyProduct", Value: WalletUserBillTypeBuyProduct},
	{Label: "购买等级", Field: "WalletUserBillTypeBuyLevel", Value: WalletUserBillTypeBuyLevel},
	{Label: "产品退款", Field: "WalletUserBillTypeRefundProduct", Value: WalletUserBillTypeRefundProduct},
	{Label: "产品收益", Field: "WalletUserBillTypeEarnings", Value: WalletUserBillTypeEarnings},
	{Label: "余额奖励", Field: "WalletUserBillTypeAward", Value: WalletUserBillTypeAward},
	{Label: "注册奖励", Field: "WalletUserBillTypeRegisterAward", Value: WalletUserBillTypeRegisterAward},
	{Label: "邀请奖励", Field: "WalletUserBillTypeShareAward", Value: WalletUserBillTypeShareAward},
	{Label: "团队收益", Field: "WalletUserBillTypeTeamEarnings", Value: WalletUserBillTypeTeamEarnings},
	{Label: "资产充值", Field: "WalletUserBillTypeAssetsDeposit", Value: WalletUserBillTypeAssetsDeposit},
	{Label: "资产系统充值", Field: "WalletUserBillTypeSystemAssetsDeposit", Value: WalletUserBillTypeSystemAssetsDeposit},
	{Label: "资产提现拒绝", Field: "WalletUserBillTypeAssetsWithdraw", Value: WalletUserBillTypeAssetsWithdraw},
	{Label: "资产系统扣款", Field: "WalletUserBillTypeSystemAssetsWithdraw", Value: WalletUserBillTypeSystemAssetsWithdraw},
	{Label: "资产提现拒绝", Field: "WalletUserBillTypeAssetsWithdrawRefuse", Value: WalletUserBillTypeAssetsWithdrawRefuse},
	{Label: "资产购买产品", Field: "WalletUserBillTypeAssetsBuyProduct", Value: WalletUserBillTypeAssetsBuyProduct},
	{Label: "资产退款产品", Field: "WalletUserBillTypeAssetsRefundProduct", Value: WalletUserBillTypeAssetsRefundProduct},
	{Label: "资产产品收益", Field: "WalletUserBillTypeAssetsEarnings", Value: WalletUserBillTypeAssetsEarnings},
	{Label: "资产奖励", Field: "WalletUserBillTypeAssetsAward", Value: WalletUserBillTypeAssetsAward},
	{Label: "团队资产收益", Field: "WalletUserBillTypeTeamAssetsEarnings", Value: WalletUserBillTypeTeamAssetsEarnings},
	{Label: "资产转换提现", Field: "WalletUserBillTypeAssetsConvertWithdraw", Value: WalletUserBillTypeAssetsConvertWithdraw},
	{Label: "资产转换充值", Field: "WalletUserBillTypeAssetsConvertDeposit", Value: WalletUserBillTypeAssetsConvertDeposit},
	{Label: "资产转移提现", Field: "WalletUserBillTypeAssetsTransferWithdraw", Value: WalletUserBillTypeAssetsTransferWithdraw},
	{Label: "资产转移充值", Field: "WalletUserBillTypeAssetsTransferDeposit", Value: WalletUserBillTypeAssetsTransferDeposit},
	{Label: "余额转移提现", Field: "WalletUserBillTypeTransferWithdraw", Value: WalletUserBillTypeTransferWithdraw},
	{Label: "余额转移充值", Field: "WalletUserBillTypeTransferDeposit", Value: WalletUserBillTypeTransferDeposit},
	{Label: "余额转换提现", Field: "WalletUserBillTypeConvertWithdraw", Value: WalletUserBillTypeConvertWithdraw},
	{Label: "余额转换充值", Field: "WalletUserBillTypeConvertDeposit", Value: WalletUserBillTypeConvertDeposit},
	{Label: "余额渠道充值", Field: "WalletUserBillTypeChannelDeposit", Value: WalletUserBillTypeChannelDeposit},
	{Label: "余额渠道提现", Field: "WalletUserBillTypeChannelWithdraw", Value: WalletUserBillTypeChannelWithdraw},
	{Label: "商家采购产品", Field: "WalletUserBillTypePurchaseProduct", Value: WalletUserBillTypePurchaseProduct},
}

// WalletUserBill 钱包账单
type WalletUserBill struct {
	types.GormModel
	AdminId  uint    `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	UserId   uint    `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	AssetsId uint    `gorm:"type:int unsigned not null;comment:资产ID" json:"assetsId"`
	SourceId uint    `gorm:"type:int unsigned not null;comment:来源ID" json:"sourceId"`
	Type     int     `gorm:"type:smallint not null;default:1;comment:类型" json:"type"`
	Name     string  `gorm:"type:varchar(60) not null;comment:名称" json:"name"`
	Money    float64 `gorm:"type:decimal(16,4) not null;comment:金额" json:"money"`
	Balance  float64 `gorm:"type:decimal(16,4) not null;comment:余额" json:"balance"`
	Data     string  `gorm:"type:text;comment:数据" json:"data"`
}
