package messages

import (
	"gofiber/app/models/model/chatsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	Key        string             `json:"key" validate:"required"` //	会话Key
	Pagination *scopes.Pagination `json:"pagination"`              //	分页
}

// Index 消息列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*chatsModel.Messages, 0)}
	database.Db.Model(&chatsModel.Messages{}).Where("`key` = ?", params.Key).
		Where("status = ?", chatsModel.MessagesStatusActive).
		Where(database.Db.Or("sender_id = ?", ctx.UserId).Or("receiver_id = ?", ctx.UserId)).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
