package product

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName  string `json:"adminName"`  // 管理
	StoreName  string `json:"storeName"`  // 店铺名
	CategoryId int    `json:"categoryId"` // 类目ID
	Name       string `json:"name"`       // 标题
	Type       int    `json:"type"`       // 类型1默认类型
	Status     int    `json:"status"`     // 状态-1禁用 10启用
	context.IndexParams
}

type product struct {
	productsModel.Product
	StoreInfo shopsModel.Store                `json:"storeInfo" gorm:"foreignKey:StoreId"`
	AdminInfo adminsModel.AdminUser           `json:"adminInfo" gorm:"foreignKey:AdminId"`
	Attrs     []productAttrsKey               `json:"attrs" gorm:"foreignKey:ProductId;references:ID"`
	SkuList   []productsModel.ProductAttrsSku `json:"skuList" gorm:"foreignKey:ProductId;references:ID"`
}

type productAttrsKey struct {
	productsModel.ProductAttrsKey
	Values []productsModel.ProductAttrsVal `json:"values" gorm:"foreignKey:KeyId;references:ID"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	// 查询params中分类ID的数据
	ids := make([]int, 0)
	if params.CategoryId > 0 {
		ids = append(ids, params.CategoryId)
		database.Db.Raw("WITH RECURSIVE cte AS (SELECT id, parent_id FROM category WHERE parent_id = ? UNION ALL  SELECT c.id, c.parent_id FROM category c JOIN cte ON cte.id = c.parent_id) SELECT id FROM cte;", params.CategoryId).
			Scan(&ids)
	}

	data := &context.IndexData{Items: make([]*product, 0)}
	database.Db.Model(&productsModel.Product{}).Preload("AdminInfo").Preload("StoreInfo").Preload("Attrs.Values").Preload("SkuList").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("store_id", &shopsModel.Store{}, "id", "name = ?", params.StoreName).
			Like("name", "%"+params.Name+"%").
			Eq("type", params.Type).
			Eq("status", params.Status).
			In("category_id", ids).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
