package walletsModel

import "gofiber/app/models/model/types"

const (
	// WalletAssetsTypeFiatCurrency 法币资产
	WalletAssetsTypeFiatCurrency = 1

	// WalletAssetsTypeDigitalCurrency 数字货币资产
	WalletAssetsTypeDigitalCurrency = 11

	// WalletAssetsTypeVirtualCurrency 虚拟货币资产
	WalletAssetsTypeVirtualCurrency = 21

	// WalletAssetsStatusActive 开启
	WalletAssetsStatusActive = 10

	// WalletAssetsStatusDisable 禁用
	WalletAssetsStatusDisable = -1
)

// WalletAssets 钱包资产管理
type WalletAssets struct {
	types.GormModel
	AdminId uint    `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	Name    string  `gorm:"type:varchar(60) not null;comment:名称" json:"name"`
	Symbol  string  `gorm:"type:varchar(60) not null;comment:标识" json:"symbol"`
	Icon    string  `gorm:"type:varchar(60) not null;comment:图标" json:"icon"`
	Type    int     `gorm:"type:tinyint not null;default:1;comment:类型 1法币资产 11数字货币资产 21虚拟货币资产" json:"type"`
	Rate    float64 `gorm:"type:decimal(16,4) not null;default:1;comment:汇率" json:"rate"`
	Status  int     `gorm:"type:smallint not null;default:10;comment:状态 -1禁用 10开启" json:"status"`
	Data    string  `gorm:"type:text;comment:数据" json:"data"`
}
