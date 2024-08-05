package userOrder

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type PresetParams struct {
	Type int //	充值类型
}

type IndexParams struct {
	AdminName string `json:"adminName"` //	管理账户
	UserName  string `json:"userName"`  //	用户账户
	AssetsId  int    `json:"assetsId"`  // 	资产ID
	OrderSn   string `json:"orderSn"`   // 	编号
	Status    int    `json:"status"`    // 	状态  -1拒绝 10审核 20同意
	context.IndexParams
}

type walletUserOrder struct {
	walletsModel.WalletUserOrder
	RateMoney   float64                        `gorm:"-"` // 汇率金额
	AdminInfo   adminsModel.AdminUser          `gorm:"foreignKey:AdminId;" json:"adminInfo"`
	UserInfo    usersModel.User                `gorm:"foreignKey:UserId;" json:"userInfo"`
	AssetsInfo  walletsModel.WalletAssets      `gorm:"foreignKey:AssetsId;" json:"assetsInfo"`
	PaymentInfo walletsModel.WalletPayment     `gorm:"foreignKey:SourceId" json:"paymentInfo"`
	AccountInfo walletsModel.WalletUserAccount `gorm:"foreignKey:SourceId" json:"accountInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	presetParams := ctx.Locals("preset").(*PresetParams)
	data := &context.IndexData{Items: make([]*walletUserOrder, 0)}
	database.Db.Model(&walletsModel.WalletUserOrder{}).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Where("type = ?", presetParams.Type).
		Preload("AdminInfo").Preload("UserInfo").Preload("AssetsInfo").Preload("PaymentInfo").Preload("AccountInfo").
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Eq("assets_id", params.AssetsId).
			Eq("order_sn", params.OrderSn).
			Eq("status", params.Status).
			Between("created_at", params.CreatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	for _, order := range data.Items.([]*walletUserOrder) {
		switch order.Type {
		case walletsModel.WalletUserOrderTypeDeposit:
			order.RateMoney = order.Money * order.PaymentInfo.Rate
		case walletsModel.WalletUserOrderTypeWithdraw:
			if order.PaymentInfo.Rate > 0 {
				order.RateMoney = order.Money / order.PaymentInfo.Rate
			}
		}
	}

	return ctx.SuccessJson(data)
}
