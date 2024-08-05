package userAccount

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` //	管理账户
	UserName  string `json:"userName"`  // 用户账户
	Mode      int    `json:"mode"`      //	模式
	PaymentId int    `json:"paymentId"` // 支付ID
	Name      string `json:"name"`      // 银行名称｜Token
	RealName  string `json:"realName"`  // 真实姓名
	Remark    string `json:"remark"`    // 备注信息
	Number    string `json:"number"`    // 卡号|地址
	Code      string `json:"code"`      // 银行代码
	Status    int    `json:"status"`    // 状态 -1禁用 10开启
	context.IndexParams
}

type walletUserAccount struct {
	walletsModel.WalletUserAccount
	AdminInfo   adminsModel.AdminUser      `gorm:"foreignKey:AdminId;" json:"adminInfo"`
	UserInfo    usersModel.User            `gorm:"foreignKey:UserId" json:"userInfo"`
	PaymentInfo walletsModel.WalletPayment `gorm:"foreignKey:PaymentId" json:"paymentInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*walletUserAccount, 0)}
	database.Db.Model(&walletsModel.WalletUserAccount{}).Preload("AdminInfo").Preload("UserInfo").Preload("PaymentInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			FindModeIn("payment_id", &walletsModel.WalletPayment{}, "id", "mode = ?", params.Mode).
			Eq("payment_id", params.PaymentId).
			Eq("name", params.Name).
			Eq("remark", params.Remark).
			Eq("real_name", params.RealName).
			Eq("number", params.Number).
			Eq("code", params.Code).
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
