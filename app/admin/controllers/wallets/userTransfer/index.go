package userTransfer

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName    string `json:"adminName"`    //	管理账户
	SenderName   string `json:"senderName"`   //	发送账户
	ReceiverName string `json:"receiverName"` //	接收账户
	Type         int    `json:"type"`         //  类型 1余额 11资产
	AssetsId     int    `json:"assetsId"`     //  资产ID
	Status       int    `json:"status"`       // 状态 -1禁用 10开启
	context.IndexParams
}

type walletUserTransfer struct {
	walletsModel.WalletUserTransfer
	AdminInfo    adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
	SenderInfo   usersModel.User       `gorm:"foreignKey:SenderId" json:"senderInfo"`
	ReceiverInfo usersModel.User       `gorm:"foreignKey:ReceiverId" json:"receiverInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*walletUserTransfer, 0)}
	database.Db.Model(&walletsModel.WalletUserTransfer{}).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Preload("AdminInfo").Preload("SenderInfo").Preload("ReceiverInfo").
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("sender_id", &usersModel.User{}, "id", "user_name = ?", params.SenderName).
			FindModeIn("receiver_id", &usersModel.User{}, "id", "user_name = ?", params.ReceiverName).
			Eq("type", params.Type).
			Eq("assets_id", params.AssetsId).
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
