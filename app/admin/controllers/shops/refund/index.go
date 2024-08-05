package refund

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理名称
	UserName  string `json:"userName"`  // 用户账户
	StoreName string `json:"storeName"` // 店铺名称
	OrderSn   string `json:"orderSn"`   // 产品订单编号
	Name      string `json:"name"`      // 申请理由
	Status    int    `json:"status"`    // 售后状态 -2删除 10申请中 11处理中 12拒绝 13完成
	context.IndexParams
}

type IndexData struct {
	shopsModel.StoreRefund
	AdminInfo adminsModel.AdminUser      `json:"adminInfo" gorm:"foreignKey:AdminId"`
	UserInfo  usersModel.User            `json:"userInfo" gorm:"foreignKey:UserId"`
	StoreInfo shopsModel.Store           `json:"storeInfo" gorm:"foreignKey:StoreId"`
	OrderInfo productsModel.ProductOrder `json:"orderInfo" gorm:"foreignKey:OrderId"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}
	database.Db.Model(&shopsModel.StoreRefund{}).
		Preload("AdminInfo").
		Preload("UserInfo").
		Preload("StoreInfo").
		Preload("OrderInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			FindModeIn("store_id", &shopsModel.Store{}, "id", "name = ?", params.StoreName).
			FindModeIn("order_id", &productsModel.ProductOrder{}, "id", "order_sn = ?", params.OrderSn).
			Eq("status", params.Status).
			Like("name", "%"+params.Name+"%").
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
