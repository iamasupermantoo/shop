package transfer

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AssetsId int // 资产ID
	Type     int // 模式
	context.IndexParams
}

type walletUserTransfer struct {
	walletsModel.WalletUserTransfer
	AssetsInfo   walletsModel.WalletAssets `json:"assetsInfo" gorm:"foreignKey:ID;references:AssetsId"`
	SenderInfo   userInfo                  `json:"senderInfo" gorm:"foreignKey:SenderId;"`
	ReceiverInfo userInfo                  `json:"receiverInfo" gorm:"foreignKey:ReceiverId;"`
}

type userInfo struct {
	ID       uint   `json:"id"`       // ID
	UserName string `json:"userName"` // 用户名
	Avatar   string `json:"avatar"`   // 头像
}

func (userInfo) TableName() string {
	return "user"
}

// Index 资金转移列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*walletUserTransfer, 0)}
	model := database.Db.Model(&walletsModel.WalletUserTransfer{}).Where("sender_id = ?", ctx.UserId).
		Preload("SenderInfo").Preload("ReceiverInfo").Preload("AssetsInfo")
	model.Scopes(scopes.NewScopes().
		Eq("assets_id", params.AssetsId).
		Eq("type", params.Type).
		Between("created_at", params.CreatedAt.ToAddDate()).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)
	return ctx.SuccessJson(data)
}
