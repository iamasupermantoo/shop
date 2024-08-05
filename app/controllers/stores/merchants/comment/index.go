package storeComment

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	Status     int                `json:"status"`     // 订单状态
	Pagination *scopes.Pagination `json:"pagination"` // 分页数据
}

// Index 评论列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: []*storeComment{}}
	storeInfo := &shopsModel.Store{}
	database.Db.Model(storeInfo).Where("user_id = ?", ctx.UserId).
		Where("status = ?", shopsModel.StoreStatusActivate).
		Find(storeInfo)
	database.Db.Model(&shopsModel.StoreComment{}).
		Where("store_id = ?", storeInfo.ID).
		Preload("StoreInfo").
		Preload("ProductInfo").
		Preload("UserInfo").
		Preload("OrderInfo").
		Scopes(scopes.NewScopes().
			Eq("status", params.Status).
			Scopes()).
		Scopes(params.Pagination.Scopes()).
		Count(&data.Count).Find(&data.Items)

	return ctx.SuccessJson(data)
}

type storeComment struct {
	shopsModel.StoreComment
	SkuInfo     productsModel.ProductAttrsSku `json:"skuInfo" gorm:"-"`
	StoreInfo   shopsModel.Store              `json:"storeInfo" gorm:"foreignKey:ID;references:StoreId"`
	UserInfo    usersModel.User               `json:"userInfo" gorm:"foreignKey:ID;references:UserId"`
	ProductInfo productsModel.Product         `json:"productInfo" gorm:"foreignKey:ID;references:ProductId"`
	OrderInfo   productsModel.ProductOrder    `json:"orderInfo" gorm:"foreignKey:ID;references:OrderId"`
}

func (storeComment) TableName() string {
	return "store_comment"
}
