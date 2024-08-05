package messages

import (
	"gofiber/app/models/model/chatsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/websocket"
	"gofiber/module/socket"
	"gofiber/utils"
)

type SendParams struct {
	Mode int    `json:"mode" validate:"required,oneof=1 2"`           // 消息类型
	Key  string `json:"key" validate:"required"`                      // 唯一标识
	Type int    `json:"type" validate:"required,oneof=1 11 12 13 20"` // 消息类型
	Data string `json:"data" validate:"required"`                     // 消息内容
}

func Send(ctx *context.CustomCtx, params *SendParams) error {
	conversationInfo := &chatsModel.Conversation{}
	result := database.Db.Model(&chatsModel.Conversation{}).Where("user_id = ?", ctx.UserId).Where("`key` = ?", params.Key).Take(conversationInfo)
	if result.Error != nil {
		return ctx.ErrorJsonTranslate("findError")
	}

	messageInfo := &chatsModel.Messages{
		AdminId:    conversationInfo.AdminId,
		Key:        conversationInfo.Key,
		SenderId:   conversationInfo.UserId,
		ReceiverId: conversationInfo.ReceiverId,
		Type:       params.Type,
		Data:       params.Data,
	}
	result = database.Db.Create(&messageInfo)
	if result.Error != nil {
		return ctx.ErrorJsonTranslate("operateDataError")
	}

	// 更新会话最新时间
	database.Db.Model(&chatsModel.Conversation{}).Where("`key` = ?", params.Key).Update("data", utils.JsonMarshal(messageInfo))

	// 发送websocket消息
	websocket.HomeWebSocket.RedisUserPublish(ctx.Rds, socket.MessageOperateMessage, messageInfo.ReceiverId, &MessageInfo{Messages: *messageInfo, Mode: params.Mode})

	return ctx.SuccessJson(messageInfo)
}

type MessageInfo struct {
	chatsModel.Messages
	Mode int `json:"mode"`
}
