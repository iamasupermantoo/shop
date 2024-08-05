package productBrowsing

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName   string `json:"adminName"`   // 管理名称
	StoreName   string `json:"storeName"`   // 商铺名称
	ProductName string `json:"productName"` // 产品名称
	Nums        int    `json:"nums"`        // 浏览次数
	Type        int    `json:"type"`        // 类型
	Status      int    `json:"status"`      // 状态
	context.IndexParams
}

type IndexData struct {
	productsModel.ProductBrowsing
	AdminInfo   adminsModel.AdminUser `json:"adminInfo" gorm:"foreignKey:AdminId"`
	StoreInfo   shopsModel.Store      `json:"storeInfo" gorm:"foreignKey:StoreId"`
	ProductInfo productsModel.Product `json:"productInfo" gorm:"foreignKey:ProductId"`
}

// Index 店铺浏览记录
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}
	//	过滤参数
	database.Db.Model(&productsModel.ProductBrowsing{}).
		Preload("StoreInfo").
		Preload("ProductInfo").
		Preload("AdminInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("store_id", &shopsModel.Store{}, "id", "name = ?", params.StoreName).
			FindModeIn("product_id", &productsModel.Product{}, "id", "name = ?", params.ProductName).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
