package comment

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
	AdminName string `json:"adminName"` // 管理账户
	UserName  string `json:"userName"`  // 用户账户
	OrderSn   string `json:"orderSn"`   // 产品订单编号
	StoreName string `json:"storeName"` // 店铺名称
	Name      string `json:"name"`      // 评论内容
	Status    int    `json:"status"`    // 评论状态
	context.IndexParams
}

type IndexData struct {
	shopsModel.StoreComment
	AdminInfo   adminsModel.AdminUser      `json:"adminInfo" gorm:"foreignKey:AdminId;"`
	UserInfo    usersModel.User            `json:"userInfo" gorm:"foreignKey:UserId;"`
	ProductInfo productsModel.Product      `json:"productInfo" gorm:"foreignKey:ProductId"`
	OrderInfo   productsModel.ProductOrder `json:"orderInfo" gorm:"foreignKey:OrderId"`
	StoreInfo   shopsModel.Store           `json:"storeInfo" gorm:"foreignKey:StoreId"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}
	database.Db.Model(&shopsModel.StoreComment{}).
		Preload("AdminInfo").
		Preload("UserInfo").
		Preload("ProductInfo").
		Preload("OrderInfo").
		Preload("StoreInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			FindModeIn("order_id", &productsModel.ProductOrder{}, "id", "order_sn = ?", params.OrderSn).
			FindModeIn("store_id", &shopsModel.Store{}, "id", "name = ?", params.StoreName).
			Like("name", "%"+params.Name+"%").
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
