package walletsModel

import "gofiber/app/models/model/types"

const (
	// WalletUserTransferStatusActive 完成
	WalletUserTransferStatusActive = 10

	// WalletUserTransferStatusDisable 失败
	WalletUserTransferStatusDisable = -1

	// WalletUserTransferTypeBalance 余额划转
	WalletUserTransferTypeBalance = 1

	// WalletUserTransferTypeAssets 资产划转
	WalletUserTransferTypeAssets = 11
)

// WalletUserTransfer 钱包资金资产转换
type WalletUserTransfer struct {
	types.GormModel
	AdminId    uint    `gorm:"type:int unsigned not null;comment: 管理员ID" json:"adminId"`
	SenderId   uint    `gorm:"type:int unsigned not null;comment:发送者ID" json:"senderId"`
	ReceiverId uint    `gorm:"type:int unsigned not null;comment:接收者ID" json:"receiverId"`
	Type       int     `gorm:"type: tinyint not null;comment:类型1余额 11资产" json:"type"`
	AssetsId   uint    `gorm:"type: int unsigned not null;comment:资产ID" json:"assetsId"`
	Money      float64 `gorm:"type:decimal(16,4) not null;comment:金额" json:"money"`
	Fee        float64 `gorm:"type: decimal(16,4) not null;comment:手续费" json:"fee"`
	Status     int     `gorm:"type:smallint not null;default:10;comment:状态 -1失败 10完成" json:"status"`
	Data       string  `gorm:"type: text;comment:数据" json:"data"`
}
