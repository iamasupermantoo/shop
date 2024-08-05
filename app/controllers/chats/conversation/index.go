package conversation

import (
	"gofiber/app/models/model/chatsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
	"gorm.io/gorm"
)

type IndexParams struct {
	Mode       int                `json:"mode" validate:"required,oneof=1 2"` // 消息类型
	Pagination *scopes.Pagination `json:"pagination"`                         // 分页
}

// Index 会话列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*Conversation, 0)}

	database.Db.Model(&chatsModel.Conversation{}).Where("user_id = ?", ctx.UserId).Where("data <> ''").Where("mode  = ?", params.Mode).
		Preload("UnreadInfo", func(db *gorm.DB) *gorm.DB {
			return db.Select("sender_id", "count(*) as Nums").Where("unread = ?", chatsModel.MessagesUnreadOff).Group("sender_id")
		}).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	for _, item := range data.Items.([]*Conversation) {
		switch params.Mode {
		case chatsModel.ConversationModeUser:
			storeInfo := shopsModel.Store{}
			database.Db.Where("user_id = ?", item.ReceiverId).Find(&storeInfo)
			item.ReceiverInfo = storeInfo

		case chatsModel.ConversationModeStore:
			userInfo := usersModel.User{}
			database.Db.Where("id = ?", item.ReceiverId).Find(&userInfo)
			item.ReceiverInfo = userInfo
		}
	}
	return ctx.SuccessJson(data)
}

type Conversation struct {
	chatsModel.Conversation
	ReceiverInfo any        `json:"receiverInfo" gorm:"-"`
	UnreadInfo   unreadInfo `json:"unreadInfo" gorm:"foreignKey:SenderId;references:ReceiverId"` //	未读消息数量
}

func (Conversation) TableName() string {
	return "conversation"
}

type unreadInfo struct {
	SenderId uint `json:"senderId"` // 接收者
	Nums     uint `json:"nums"`     // 未读消息数
}

func (unreadInfo) TableName() string {
	return "messages"
}
