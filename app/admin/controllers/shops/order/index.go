package shopsOrder

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"`
	UserName  string `json:"userName"`
	StoreName string `json:"storeName"`
	OrderSn   string `json:"orderSn"` // 店铺订单编号
	Status    int    `json:"status"`  // -1取消 10待支付 12待发货 14待收货 20完成
	context.IndexParams
}

type IndexData struct {
	shopsModel.ProductStoreOrder
	AdminInfo adminsModel.AdminUser `json:"adminInfo" gorm:"foreignKey:AdminId"`
	UserInfo  usersModel.User       `json:"userInfo" gorm:"foreignKey:UserId"`
	StoreInfo shopsModel.Store      `json:"storeInfo" gorm:"foreignKey:StoreId"`
}

// Index 店铺订单列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}
	//	过滤参数
	database.Db.Model(&shopsModel.ProductStoreOrder{}).
		Preload("StoreInfo").
		Preload("AdminInfo").
		Preload("UserInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			FindModeIn("store_id", &shopsModel.Store{}, "id", "name = ?", params.StoreName).
			Eq("order_sn", params.OrderSn).
			Eq("status", params.Status).
			Between("updated_at BETWEEN ? AND ?", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
