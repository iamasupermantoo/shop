package chatsModel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"gofiber/app/models/model/types"
)

const (
	// ConversationStatusActive 会话状态正常
	ConversationStatusActive = 10

	// ConversationStatusDisable 会话状态结束
	ConversationStatusDisable = -1

	// ConversationTypePrivate 私信
	ConversationTypePrivate = 1

	// ConversationOnline 在线
	ConversationOnline = 1

	// ConversationOffline 离线
	ConversationOffline = 2

	// ConversationModeUser 用户会话模式
	ConversationModeUser = 1

	// ConversationModeStore 商家会话模式
	ConversationModeStore = 2
)

// Conversation 聊天会话
type Conversation struct {
	types.GormModel
	Key        string       `json:"key" gorm:"type:varchar(255) not null;comment:会话Key"`
	AdminId    uint         `json:"adminId" gorm:"type:int unsigned not null;comment:管理ID"`
	UserId     uint         `json:"userId" gorm:"type:int unsigned not null;comment:用户ID"`
	ReceiverId uint         `json:"receiverId" gorm:"type:int unsigned not null;comment:接收ID"`
	Name       string       `json:"name" gorm:"type:varchar(120) not null;comment:备注"`
	Type       int          `json:"type" gorm:"type:tinyint not null default 1;default:1;comment:类型 1私信"`
	Mode       int          `json:"mode" gorm:"type:tinyint not null default 1;default:1;comment:会话模式 1用户 2商家"`
	Online     int          `json:"online" gorm:"type:tinyint(1);default:1;comment:在线 1离线 2在线"`
	Status     int          `json:"status" gorm:"type:smallint not null default 10;default:10;comment:状态-1屏蔽 10正常"`
	Data       *MessageData `json:"data" gorm:"type:text;comment:数据"`
}

// MessageData 	会话数据
type MessageData Messages

func (_MessageData *MessageData) Scan(val interface{}) error {
	bytes, ok := val.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan MessageData value:", val))
	}
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, _MessageData)
	}
	*_MessageData = MessageData{}

	return nil
}

func (_MessageData *MessageData) Value() (driver.Value, error) {
	if _MessageData == nil {
		return json.Marshal(&MessageData{})
	}
	return json.Marshal(_MessageData)
}
