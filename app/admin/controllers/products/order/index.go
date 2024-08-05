package order

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理账户
	UserName  string `json:"userName"`  // 用户账户
	OrderSn   string `json:"orderSn"`   // 订单编号
	ProductId int    `json:"productId"` // 产品ID
	Type      int    `json:"type"`      // 订单类型
	Status    int    `json:"status"`    //  订单状态 -1取消 10等待 11运行 20完成
	context.IndexParams
}

type productOrder struct {
	productsModel.ProductOrder
	AdminInfo   adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
	UserInfo    usersModel.User       `gorm:"foreignKey:UserId;" json:"userInfo"`
	ProductInfo productsModel.Product `gorm:"foreignKey:ProductId;" json:"productInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*productOrder, 0)}
	database.Db.Model(&productsModel.ProductOrder{}).
		Preload("AdminInfo").Preload("UserInfo").Preload("ProductInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Eq("order_sn", params.OrderSn).
			Eq("product_id", params.ProductId).
			Eq("type", params.Type).
			Eq("status", params.Status).
			Between("created_at", params.CreatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
