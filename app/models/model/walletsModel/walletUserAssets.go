package walletsModel

import "gofiber/app/models/model/types"

const (
	// WalletUserAssetsStatusActive 开启
	WalletUserAssetsStatusActive = 10

	// WalletUserAssetsStatusDisable 禁用
	WalletUserAssetsStatusDisable = -1
)

// WalletUserAssets 用户钱包资产
type WalletUserAssets struct {
	types.GormModel
	AdminId  uint    `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	UserId   uint    `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	AssetsId uint    `gorm:"type:int unsigned not null;comment:资产ID" json:"assetsId"`
	Money    float64 `gorm:"type:decimal(16,4) not null;comment:金额" json:"money"`
	Status   int     `gorm:"type:smallint not null;default:10;comment:状态 -1禁用 10开启" json:"status"`
	Data     string  `gorm:"type:text;comment:数据" json:"data"`
}
