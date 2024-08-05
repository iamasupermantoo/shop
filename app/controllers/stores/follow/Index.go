package follow

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	Type       int                `json:"type" validate:"required"` // 收藏类型  1收藏店铺 2收藏产品
	Pagination *scopes.Pagination `json:"pagination"`               // 分页数据
}

// Index 收藏产品列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := context.IndexData{Items: make([]*followInfo, 0)}
	database.Db.Model(&shopsModel.StoreFollow{}).
		Preload("ProductInfo").Preload("StoreInfo").
		Where("status = ?", shopsModel.StoreFollowStatusConcern).
		Where("user_id = ?", ctx.UserId).
		Scopes(scopes.NewScopes().
			Eq("type", params.Type).Scopes()).
		Scopes(params.Pagination.Scopes()).
		Count(&data.Count).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}

// followInfo 关注列表
type followInfo struct {
	shopsModel.StoreFollow
	ProductInfo productsModel.Product `json:"productInfo" gorm:"foreignKey:ProductId;"`
	StoreInfo   shopsModel.Store      `json:"storeInfo" gorm:"foreignKey:StoreId;"`
}

func (_followInfo *followInfo) TableName() string {
	return "store_follow"
}
