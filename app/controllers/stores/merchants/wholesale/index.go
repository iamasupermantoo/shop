package wholesale

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	Status     int                `json:"status"`     // 商品状态			0批发商品 -1店铺商品下架 10店铺商品上架
	Search     string             `json:"search"`     // 搜索商品详情内容
	Pagination *scopes.Pagination `json:"pagination"` // 分页数据
}

type wholesaleProduct struct {
	productsModel.Product
	SkuList []productsModel.ProductAttrsSku `json:"skuList" gorm:"foreignKey:ProductId"`
}

func (wholesaleProduct) TableName() string {
	return "product"
}

func Index(ctx *context.CustomCtx, params *IndexParams) error {
	storeInfo := &shopsModel.Store{}
	result := database.Db.Model(storeInfo).
		Where("user_id = ?", ctx.UserId).
		Where("status = ?", shopsModel.StoreStatusActivate).
		Find(storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	data := context.IndexData{Items: make([]*wholesaleProduct, 0)}

	model := database.Db.Model(&productsModel.Product{}).
		Where("admin_id = ?", ctx.AdminSettingId)
	if params.Status == 0 {
		// 批发商品
		var ids []int
		database.Db.Raw("select parent_id from product where store_id = ?", storeInfo.ID).Scan(&ids)
		model.Where("type = ?", productsModel.ProductTypeWholesale)
		if len(ids) > 0 {
			model.Where("id NOT IN ?", ids)
		}
	} else {
		// 店铺商品
		model.Where("store_id = ?", storeInfo.ID).
			Where("status = ?", params.Status)
	}
	model.Preload("SkuList").
		Scopes(scopes.NewScopes().
			Like("name", params.Search+"%").Scopes()).
		Scopes(params.Pagination.Scopes()).
		Count(&data.Count).
		Find(&data.Items)

	// 如果批发商品, 那么把关联的店铺产品状态赋值出去
	for _, product := range data.Items.([]*wholesaleProduct) {
		if product.SkuList == nil {
			product.SkuList = make([]productsModel.ProductAttrsSku, 0)
		}
	}

	return ctx.SuccessJson(data)
}
