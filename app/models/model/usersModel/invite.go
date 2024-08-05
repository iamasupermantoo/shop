package usersModel

import "gofiber/app/models/model/types"

const (
	// InviteTypeAdmin 管理邀请码
	InviteTypeAdmin = 1

	// InviteTypeUser 用户邀请码
	InviteTypeUser = 2

	// InviteStatusActive 开启
	InviteStatusActive = 10

	// InviteStatusDisable 禁用
	InviteStatusDisable = -1
)

// Invite 用户邀请码
type Invite struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	UserId  uint   `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	Code    string `gorm:"type:varchar(50) not null;uniqueIndex;comment:邀请码" json:"code"`
	Status  int    `gorm:"type:tinyint not null;default:10;comment:状态 -1禁用 10开启" json:"status"`
	Data    string `gorm:"type:text;comment:数据" json:"data"`
}
