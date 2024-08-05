package walletsModel

import "gofiber/app/models/model/types"

const (
	// WalletUserAccountStatusActive 开启
	WalletUserAccountStatusActive = 10

	// WalletUserAccountStatusDisable 禁用
	WalletUserAccountStatusDisable = -1
)

// WalletUserAccount 钱包卡片管理
type WalletUserAccount struct {
	types.GormModel
	AdminId   uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	UserId    uint   `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	PaymentId uint   `gorm:"type:int unsigned not null;comment:支付ID" json:"paymentId"`
	Name      string `gorm:"type:varchar(50) not null;comment:银行名称｜Token" json:"name"`
	RealName  string `gorm:"type:varchar(50) not null;comment:真实姓名" json:"realName"`
	Number    string `gorm:"type:varchar(50) not null;comment:卡号|地址" json:"number"`
	Code      string `gorm:"type:varchar(255) not null;comment:银行代码" json:"code"`
	Remark    string `gorm:"type:varchar(255) not null;comment:备注信息" json:"remark"`
	Status    int    `gorm:"type:smallint not null;default:10;comment:状态 -1禁用 10开启" json:"status"`
	Data      string `gorm:"type:text;comment:数据" json:"data"`
}
