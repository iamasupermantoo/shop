package conversation

import (
	"github.com/google/uuid"
	"gofiber/app/models/model/chatsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type DetailsParams struct {
	Mode   int  `json:"mode" validate:"required,oneof=1 2"` // 消息类型
	UserId uint `json:"userId" validate:"required"`         //	用户ID
}

// Details 会话详情
func Details(ctx *context.CustomCtx, params *DetailsParams) error {
	// 用户本身不能聊天
	if ctx.UserId == params.UserId {
		return ctx.ErrorJsonTranslate("findError")
	}

	userInfo := &usersModel.User{}
	database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Find(userInfo)

	// 获取聊天对象信息
	receiverInfo := &usersModel.UserInfo{}
	result := database.Db.Model(&usersModel.User{}).Where("id = ?", params.UserId).Where("status = ?", usersModel.UserStatusActive).Find(receiverInfo)
	if result.Error != nil || receiverInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	// 查询会话信息是否存在
	conversationReceiver := &chatsModel.Conversation{}
	database.Db.Model(conversationReceiver).Where("user_id = ?", receiverInfo.ID).Where("receiver_id = ?", userInfo.ID).Find(conversationReceiver)

	conversationUser := &chatsModel.Conversation{}
	database.Db.Model(conversationUser).Where("user_id = ?", userInfo.ID).Where("receiver_id = ?", receiverInfo.ID).Find(conversationUser)

	conversationKey := uuid.New().String()
	if conversationReceiver.ID > 0 {
		conversationKey = conversationReceiver.Key
	}
	if conversationUser.ID > 0 {
		conversationKey = conversationUser.Key
	}

	if conversationReceiver.ID == 0 {
		conversationKey = uuid.New().String()
		database.Db.Create(&chatsModel.Conversation{
			Key:        conversationKey,
			AdminId:    receiverInfo.AdminId,
			UserId:     receiverInfo.ID,
			ReceiverId: userInfo.ID,
			Mode:       chatsModel.ConversationModeStore,
			Name:       userInfo.NickName,
		})
	}

	if conversationUser.ID == 0 {
		database.Db.Create(&chatsModel.Conversation{
			Key:        conversationKey,
			AdminId:    userInfo.AdminId,
			UserId:     userInfo.ID,
			ReceiverId: receiverInfo.ID,
			Mode:       chatsModel.ConversationModeUser,
			Name:       receiverInfo.NickName,
		})
	}

	return ctx.SuccessJson(&conversationData{
		Key:          conversationKey,
		ReceiverInfo: receiverInfo,
		UserInfo:     userInfo,
	})
}

type conversationData struct {
	Key          string               `json:"key"`          // 会话Key
	ReceiverInfo *usersModel.UserInfo `json:"receiverInfo"` // 对方信息
	UserInfo     *usersModel.User     `json:"userInfo"`     // 客户信息
}
