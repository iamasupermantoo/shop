package types

import (
	"gorm.io/gorm"
	"time"
)

const (
	// WalletsModeBalance 余额模式
	WalletsModeBalance = 1
	// WalletsModeAssets 资产模式
	WalletsModeAssets = 2

	// ModelBoolTrue 真
	ModelBoolTrue = 1
	// ModelBoolFalse 假
	ModelBoolFalse = 2
)

// GormModel Gorm 通用字段
type GormModel struct {
	ID        uint           `gorm:"primarykey;comment:ID" json:"id"`
	CreatedAt time.Time      `gorm:"comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
