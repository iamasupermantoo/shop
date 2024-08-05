package index

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Home 首页信息
func Home(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	settingService := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId)
	siteInfo := settingService.GetSiteInfo()
	settingAdminId := ctx.AdminSettingId

	// 推荐分类
	categoryList := make([]*productsModel.Category, 0)
	database.Db.Model(&productsModel.Category{}).
		Where("parent_id = ?", 0).
		Where("status = ?", productsModel.CategoryStatusActive).
		Where("admin_id = ?", settingAdminId).
		Limit(12).Order("sort DESC").
		Find(&categoryList)

	// 热门评分最优店铺
	ratingStoreList := make([]*shopsModel.Store, 0)
	database.Db.Model(&shopsModel.Store{}).
		Where("status = ?", shopsModel.StoreStatusActivate).
		Where("admin_id = ?", settingAdminId).
		Order("Rating DESC").
		Limit(10).
		Find(&ratingStoreList)

	// 热门商品,销量降序前50个
	salesProductList := make([]*productsModel.Product, 0)
	database.Db.Model(&productsModel.Product{}).Where("type = ?", productsModel.ProductTypeDefault).Where("admin_id = ?", settingAdminId).
		Where("status = ?", productsModel.ProductStatusActive).
		Order("rand()").
		Limit(50).
		Find(&salesProductList)

	// 每日上新
	newProductList := make([]*productsModel.Product, 0)
	database.Db.Model(productsModel.Product{}).Where("type = ?", productsModel.ProductTypeDefault).Where("admin_id = ?", settingAdminId).
		Where("status = ?", productsModel.ProductStatusActive).
		Order("created_at DESC").
		Limit(20).
		Find(&newProductList)

	return ctx.SuccessJson(&homeData{
		Notice:           siteInfo.Notice,
		Introduce:        siteInfo.Introduce,
		Banner:           settingService.GetBanner(),
		CategoryList:     categoryList,
		NewProductList:   newProductList,
		RatingStoreList:  ratingStoreList,
		SalesProductList: salesProductList,
	})
}

type homeData struct {
	Notice           string                    `json:"notice"`           // 站点公告
	Introduce        string                    `json:"introduce"`        // 站点介绍
	Banner           []string                  `json:"banner"`           // 轮播图
	CategoryList     []*productsModel.Category `json:"categoryList"`     // 产品分类
	NewProductList   []*productsModel.Product  `json:"newProductList"`   // 最新产品
	RatingStoreList  []*shopsModel.Store       `json:"ratingStoreList"`  // 评分商店
	SalesProductList []*productsModel.Product  `json:"salesProductList"` // 热门商品
}
