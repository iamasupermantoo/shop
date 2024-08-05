package messages

import (
	"gofiber/app/models/model/chatsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/websocket"
	"gofiber/module/socket"
)

type ReadParams struct {
	Ids []uint `json:"ids"` // 消息ID组
}

// Read 消息已读
func Read(ctx *context.CustomCtx, params *ReadParams) error {
	for _, id := range params.Ids {
		messageInfo := &chatsModel.Messages{}
		database.Db.Model(messageInfo).Where("id = ?", id).Where("receiver_id = ?", ctx.UserId).Find(messageInfo)
		if messageInfo.ID > 0 {
			database.Db.Model(&chatsModel.Messages{}).Where("id = ?", messageInfo.ID).Update("unread", chatsModel.MessagesUnreadOn)
			websocket.HomeWebSocket.RedisUserPublish(ctx.Rds, socket.MessageOperateReadMsg, messageInfo.SenderId, messageInfo.ID)
		}
	}

	return ctx.SuccessJsonOK()
}
