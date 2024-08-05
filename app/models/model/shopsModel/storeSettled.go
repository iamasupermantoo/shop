package shopsModel

import (
	"gofiber/app/models/model/types"
)

const (
	StoreSettledStatusRefuse  = -1 // 拒绝
	StoreSettledStatusPending = 10 // 审核中
	StoreSettledStatusPass    = 20 // 通过

	// StoreSettledTypeIdCard 身份证
	StoreSettledTypeIdCard = 1

	// StoreSettledTypeLicense 营业执照
	StoreSettledTypeLicense = 2
)

// StoreSettled 店铺入驻申请
type StoreSettled struct {
	types.GormModel
	AdminId   uint   `json:"adminId" gorm:"type:int unsigned not null;default:0;comment:管理员ID"`
	UserId    uint   `json:"userId" gorm:"type:int unsigned not null;default:0;comment:用户ID"`
	CountryId uint   `json:"countryId" gorm:"type:int unsigned not null;default:0;comment:国家ID"`
	Type      int    `json:"type" gorm:"type:tinyint not null;default:1;comment:类型 1营业执照"`
	Name      string `json:"name" gorm:"type:varchar(20) not null;default:'';comment:证件名字"`
	Address   string `json:"address" gorm:"type:varchar(2048) not null;default:'';comment:地址"`
	RealName  string `json:"realName" gorm:"type:varchar(50) not null;comment:真实姓名"`
	Logo      string `json:"logo" gorm:"type:varchar(512) not null;default:'';comment:店铺Logo"`
	Photo1    string `json:"photo1" gorm:"type:varchar(120) not null;comment:证件照1"`
	Photo2    string `json:"photo2" gorm:"type:varchar(120) not null;comment:证件照2"`
	Photo3    string `json:"photo3" gorm:"type:varchar(120) not null;comment:证件照3"`
	Number    string `json:"number" gorm:"type:varchar(255) not null;default:'';comment:证件号"`
	Email     string `json:"email" gorm:"type:varchar(50) not null;default:'';comment:email"`
	Contact   string `json:"contact" gorm:"type:varchar(50) not null;default:'';comment:联系方式"`
	Data      string `json:"data" gorm:"type:varchar(255) not null;default:'';comment:数据"`
	Status    int    `json:"status" gorm:"type:tinyint not null;default:10;comment:状态 -1拒绝 10审核中 20 通过"`
}
