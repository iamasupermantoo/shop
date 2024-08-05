package usersModel

import "gofiber/app/models/model/types"

const (
	// ChannelModeApprove 授权
	ChannelModeApprove = 1
	// ChannelModeChannel 渠道
	ChannelModeChannel = 11

	// ChannelStatusActive 激活
	ChannelStatusActive = 10
	// ChannelStatusDisable 禁用
	ChannelStatusDisable = -1

	// ChannelTypeDefault 默认类型
	ChannelTypeDefault = 1
)

// Channel 渠道表
type Channel struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;default:1;comment:管理ID" json:"adminId"`
	Name    string `gorm:"type:varchar(60) not null;comment:名称" json:"name"`
	Symbol  string `gorm:"type:varchar(60) not null;comment:标识" json:"symbol"`
	Route   string `gorm:"type:varchar(255) not null;comment:路由" json:"route"`
	Pass    string `gorm:"type:varchar(255) not null;comment:Key" json:"pass"`
	Mode    int    `gorm:"type:tinyint not null;default:1;comment:模式1授权 11渠道" json:"mode"`
	Type    int    `gorm:"type:tinyint not null;default:1;comment:类型1默认类型" json:"type"`
	Status  int    `gorm:"type:smallint not null;default:10;comment:状态 -1禁用 10激活" json:"status"`
	Data    string `gorm:"type:varchar(255) not null;comment:数据" json:"data"`
}
