package walletsModel

import "gofiber/app/models/model/types"

const (
	// WalletUserConvertStatusActive 开启
	WalletUserConvertStatusActive = 10

	// WalletUserConvertStatusDisable 禁用
	WalletUserConvertStatusDisable = -1

	// WalletUserConvertTypeFlash 闪兑
	WalletUserConvertTypeFlash = 11
)

// WalletUserConvert 钱包资金资产转换
type WalletUserConvert struct {
	types.GormModel
	AdminId    uint    `gorm:"type:int unsigned not null;comment:管理员ID" json:"adminId"`
	UserId     uint    `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	Type       int     `gorm:"type:tinyint not null;comment:类型 1闪兑" json:"type"`
	AssetsId   uint    `gorm:"type:int unsigned not null;comment:资产ID" json:"assetsId"`
	ToAssetsId uint    `gorm:"type:int unsigned not null;comment:接收资产ID" json:"toAssetsId"`
	Money      float64 `gorm:"type:decimal(16,4) not null;comment:金额" json:"money"`
	Nums       float64 `gorm:"type:decimal(16,4) not null;comment:数量" json:"nums"`
	Fee        float64 `gorm:"type:decimal(16,4) not null;comment:手续费" json:"fee"`
	Status     int     `gorm:"type:smallint not null;default:10;comment:状态 -1失败 10完成" json:"status"`
	Data       string  `gorm:"type:text;comment:数据" json:"data"`
}
