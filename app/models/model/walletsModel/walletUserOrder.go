package walletsModel

import (
	"gofiber/app/models/model/types"
	"time"
)

const (
	// WalletUserOrderStatusComplete 完成
	WalletUserOrderStatusComplete = 20

	// WalletUserOrderStatusActive 审核
	WalletUserOrderStatusActive = 10

	// WalletUserOrderStatusRefuse 拒绝
	WalletUserOrderStatusRefuse = -1

	// WalletUserOrderTypeDeposit 充值类型
	WalletUserOrderTypeDeposit = 1

	// WalletUserOrderTypeWithdraw 提现类型
	WalletUserOrderTypeWithdraw = 11

	// WalletUserOrderTypeAssetsDeposit 资产充值类型
	WalletUserOrderTypeAssetsDeposit = 2

	// WalletUserOrderTypeAssetsWithdraw 资产提现类型
	WalletUserOrderTypeAssetsWithdraw = 12
)

// WalletUserOrder 钱包订单
type WalletUserOrder struct {
	types.GormModel
	AdminId  uint    `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	UserId   uint    `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	AssetsId uint    `gorm:"type int unsigned not null;comment:资产ID" json:"assetsId"`
	SourceId uint    `gorm:"type:int unsigned not null;comment:来源ID" json:"sourceId"`
	Type     int     `gorm:"type:tinyint not null;default:1;comment:类型 1充值类型 2资产充值类型 11提现类型 12资产提现类型" json:"type"`
	OrderSn  string  `gorm:"type:varchar(60) not null;uniqueIndex;comment:编号" json:"orderSn"`
	Money    float64 `gorm:"type:decimal(16,4) not null;comment:金额" json:"money"`
	Fee      float64 `gorm:"type:decimal(8,4) not null;comment:手续费" json:"fee"`
	Voucher  string  `gorm:"type:varchar(250) not null;comment:支付凭证" json:"voucher"`
	Status   int     `gorm:"type:smallint not null;default:10;comment:状态 -1拒绝 10审核 20完成" json:"status"`
	Data     string  `gorm:"type:text;comment:数据" json:"data"`
}

// WalletUserOrderInfo 钱包订单信息
type WalletUserOrderInfo struct {
	ID          uint              `json:"id"`       // ID
	OrderSn     string            `json:"orderSn"`  // 订单号
	AssetsId    uint              `json:"assetsId"` // 资产ID
	SourceId    uint              `json:"sourceId"` // 来源ID
	Type        int               `json:"type"`     // 类型
	Money       float64           `json:"money"`    // 金额
	Fee         float64           `json:"fee"`      // 手续费
	Status      int               `json:"status"`   // 状态
	Data        string            `json:"data"`     // 拒绝理由
	PaymentInfo WalletPayment     `gorm:"foreignKey:SourceId" json:"paymentInfo"`
	AccountInfo WalletUserAccount `gorm:"foreignKey:SourceId" json:"accountInfo"`
	CreatedAt   time.Time         `json:"createdAt"` // 提交时间
}
