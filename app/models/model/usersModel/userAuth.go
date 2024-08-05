package usersModel

import "gofiber/app/models/model/types"

const (
	// UserAuthTypeIdCard 身份证
	UserAuthTypeIdCard = 1

	// UserAuthStatusComplete 完成
	UserAuthStatusComplete = 20

	// UserAuthStatusActive 审核
	UserAuthStatusActive = 10

	// UserAuthStatusRefuse 拒绝
	UserAuthStatusRefuse = -1
)

// UserAuth 实名认证
type UserAuth struct {
	types.GormModel
	AdminId  uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	UserId   uint   `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	RealName string `gorm:"type:varchar(50) not null;comment:真实姓名" json:"realName"`
	Number   string `gorm:"type:varchar(50) not null;comment:卡号" json:"number"`
	Photo1   string `gorm:"type:varchar(120) not null;comment:证件照1" json:"photo1"`
	Photo2   string `gorm:"type:varchar(120) not null;comment:证件照2" json:"photo2"`
	Photo3   string `gorm:"type:varchar(120) not null;comment:证件照3" json:"photo3"`
	Address  string `gorm:"type:varchar(255) not null;default:'';comment:详细地址" json:"address"`
	Type     int    `gorm:"type:tinyint not null;default:1;comment:类型 1身份证" json:"type"`
	Status   int    `gorm:"type:tinyint not null;default:10;comment:状态 -1拒绝 10审核 20完成" json:"status"`
	Data     string `gorm:"type:varchar(255) not null;comment:数据" json:"data"`
}
