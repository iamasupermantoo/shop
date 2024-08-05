package chatsModel

import (
	"gofiber/app/models/model/types"
)

const (
	// MessagesStatusActive 会话消息状态正常
	MessagesStatusActive = 10

	// MessagesStatusDisable 会话消息状态撤回
	MessagesStatusDisable = -1

	// MessagesUnreadOff 会话消息未读
	MessagesUnreadOff = 1

	// MessagesUnreadOn 会话消息已读
	MessagesUnreadOn = 2

	// MessagesTypeText 会话消息类型文本
	MessagesTypeText = 1

	// MessagesImage 会话消息类型图片
	MessagesImage = 11

	// MessagesAudio 会话消息类型音频
	MessagesAudio = 12

	// MessagesVideo 会话消息类型视频
	MessagesVideo = 13

	// MessagesEditor 会话消息富文本
	MessagesEditor = 20

	// MessageModeUser 用户会话消息
	MessageModeUser = 1

	// MessageModeStore 商家会话消息
	MessageModeStore = 2
)

// Messages 会话消息
type Messages struct {
	types.GormModel
	Key        string `json:"key" gorm:"type:varchar(255) not null;comment:会话Key"`
	AdminId    uint   `json:"adminId" gorm:"type:int unsigned not null;comment:管理ID"`
	SenderId   uint   `json:"senderId" gorm:"type:int unsigned not null;comment:发送ID"`
	ReceiverId uint   `json:"receiverId" gorm:"type:int unsigned not null;comment:接收ID"`
	Unread     int    `json:"unread" gorm:"type:tinyint not null default 1;default:1;comment:类型 1未读 2已读"`
	Type       int    `json:"type" gorm:"type:tinyint not null default 1;default:1;comment:类型 1文本 11图片 12语音 13视频 20富文本"`
	Status     int    `json:"status" gorm:"type:smallint not null default 10;default:10;comment:状态-1撤回 10正常"`
	Data       string `json:"data" gorm:"type:text;comment:数据"`
}
