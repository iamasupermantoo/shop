package systemsModel

import "gofiber/app/models/model/types"

const (
	// NotifyModeAdminMessage 管理通知消息
	NotifyModeAdminMessage = 1

	// NotifyModeHomeMessage 用户通知消息
	NotifyModeHomeMessage = 11

	// NotifyStatusActive 未读消息
	NotifyStatusActive = 10

	// NotifyStatusComplete 已读消息
	NotifyStatusComplete = 20
)

// Notify 系统通知表
type Notify struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	UserId  uint   `gorm:"type:int unsigned not null;comment:用户ID" json:"userId"`
	Mode    int    `gorm:"type:smallint not null;comment:模式 1后台 11前台" json:"mode"`
	Type    int    `gorm:"type:smallint not null;default:1;comment:类型1系统通知" json:"type"`
	Name    string `gorm:"type:varchar(60) not null;comment:标题" json:"name"`
	Content string `gorm:"type:text;comment:内容" json:"content"`
	Status  int    `gorm:"type:smallint not null;default:10;comment:状态" json:"status"`
	Data    string `gorm:"type:text;comment:数据" json:"data"`
}
