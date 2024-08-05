package cart

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
	AdminName   string `json:"adminName"`   // 管理账户
	UserName    string `json:"userName"`    // 用户账户
	StoreName   string `json:"storeName"`   // 店铺名称
	ProductName string `json:"productName"` // 商品名称
	context.IndexParams
}

type IndexData struct {
	shopsModel.StoreCart
	AdminInfo   adminsModel.AdminUser         `json:"adminInfo" gorm:"foreignKey:AdminId;"`
	UserInfo    usersModel.User               `json:"userInfo" gorm:"foreignKey:UserId;"`
	StoreInfo   shopsModel.Store              `json:"storeInfo" gorm:"foreignKey:StoreId;"`
	ProductInfo productsModel.Product         `json:"productInfo" gorm:"foreignKey:ProductId;"`
	SkuInfo     productsModel.ProductAttrsSku `json:"skuInfo" gorm:"foreignKey:SkuId;"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}

	//	过滤参数
	database.Db.Model(&shopsModel.StoreCart{}).
		Preload("AdminInfo").
		Preload("UserInfo").
		Preload("StoreInfo").
		Preload("ProductInfo").
		Preload("SkuInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			FindModeIn("store_id", &shopsModel.Store{}, "id", "name = ?", params.StoreName).
			FindModeIn("product_id", &productsModel.Product{}, "id", "name = ?", params.ProductName).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
