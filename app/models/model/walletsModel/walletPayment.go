package walletsModel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"gofiber/app/models/model/types"
)

const (
	// WalletPaymentStatusActive 开启
	WalletPaymentStatusActive = 10

	// WalletPaymentStatusDisable 禁用
	WalletPaymentStatusDisable = -1

	// WalletPaymentTypeBank 银行卡类型
	WalletPaymentTypeBank = 1

	// WalletPaymentTypeDigital 数字货币类型
	WalletPaymentTypeDigital = 11

	// WalletPaymentTypeChannel 渠道类型
	WalletPaymentTypeChannel = 20

	// WalletPaymentTypeThree 三方支付
	WalletPaymentTypeThree = 21

	// WalletPaymentModeDeposit 充值模式
	WalletPaymentModeDeposit = 1

	// WalletPaymentModeAssetsDeposit 资产充值模式
	WalletPaymentModeAssetsDeposit = 2

	// WalletPaymentModeWithdraw 提现模式
	WalletPaymentModeWithdraw = 11

	// WalletPaymentModeAssetsWithdraw 资产提现模式
	WalletPaymentModeAssetsWithdraw = 12
)

// WalletPayment 钱包支付管理
type WalletPayment struct {
	types.GormModel
	AdminId   uint                  `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	AssetsId  uint                  `gorm:"type:int unsigned not null;comment:default:0;钱包资产ID" json:"assetsId"`
	Name      string                `gorm:"type:varchar(60) not null;comment:名称" json:"name"`
	Icon      string                `gorm:"type:varchar(60) not null;comment:图标" json:"icon"`
	Symbol    string                `gorm:"type:varchar(60) not null;comment:标识" json:"symbol"`
	Type      int                   `gorm:"type:tinyint not null;default:1;comment:类型 1银行卡类型 11数字货币类型 21第三方支付" json:"type"`
	Mode      int                   `gorm:"type:tinyint not null;default:1;comment:模式 1充值模式 2资产充值模式 11提现模式 12资产提现模式" json:"mode"`
	Rate      float64               `gorm:"type:decimal(16,4) not null;default:1;comment:汇率" json:"rate"`
	IsVoucher int                   `gorm:"type:tinyint(1) not null;default:1;comment:显示凭证" json:"isVoucher"`
	Level     int                   `gorm:"type:tinyint not null;default:1;comment:等级" json:"level"`
	Status    int                   `gorm:"type:smallint not null;default:10;comment:状态 -1禁用 10开启" json:"status"`
	Data      GormWalletPaymentData `gorm:"type:text;comment:数据" json:"data"`
	Desc      string                `gorm:"type:text;comment:详情" json:"desc"`
}

// WalletPaymentInfo 钱包支付信息
type WalletPaymentInfo struct {
	ID        uint                  `json:"id"`        // ID
	AssetsId  uint                  `json:"assetsId"`  // 资产ID
	Symbol    string                `json:"symbol"`    // 标识
	Name      string                `json:"name"`      // 名称
	Icon      string                `json:"icon"`      // 图标
	Type      int                   `json:"type"`      // 类型
	Rate      float64               `json:"rate"`      // 汇率
	IsVoucher int                   `json:"isVoucher"` // 显示凭证
	Desc      string                `json:"desc"`      // 说明
	Data      GormWalletPaymentData `json:"data"`      // 数据
}

type GormWalletPaymentData []*WalletPaymentData

func (_GormWalletPaymentData GormWalletPaymentData) Value() (driver.Value, error) {
	if _GormWalletPaymentData == nil {
		return []byte("[]"), nil
	}
	return json.Marshal(_GormWalletPaymentData)
}

func (_GormWalletPaymentData *GormWalletPaymentData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan GormWalletPaymentData value:", value))
	}

	if len(bytes) > 0 {
		return json.Unmarshal(bytes, _GormWalletPaymentData)
	}
	*_GormWalletPaymentData = make([]*WalletPaymentData, 0)
	return nil
}

// WalletPaymentData 支付数据
type WalletPaymentData struct {
	Label  string `json:"label"`  // 名称
	Field  string `json:"field"`  // 字段
	Value  string `json:"value"`  // 内容
	IsShow bool   `json:"isShow"` // 隐藏
}
