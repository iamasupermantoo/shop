package messages

import (
	"gofiber/app/models/model/chatsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/websocket"
	"gofiber/module/socket"
)

type CancelParams struct {
	ID uint `json:"id" validate:"required"` //	消息ID
}

// Cancel 消息撤回
func Cancel(ctx *context.CustomCtx, params *CancelParams) error {
	messageInfo := &chatsModel.Messages{}
	database.Db.Model(messageInfo).Where("sender_id = ?", ctx.UserId).Where("id = ?", params.ID).Find(messageInfo)
	if messageInfo.ID > 0 {
		database.Db.Model(&chatsModel.Messages{}).Where("id = ?", messageInfo.ID).Update("status", chatsModel.MessagesStatusDisable)
		websocket.HomeWebSocket.RedisUserPublish(ctx.Rds, socket.MessageOperateCancel, messageInfo.SenderId, messageInfo.ID)
		websocket.HomeWebSocket.RedisUserPublish(ctx.Rds, socket.MessageOperateCancel, messageInfo.ReceiverId, messageInfo.ID)
	}

	return ctx.SuccessJsonOK()
}
