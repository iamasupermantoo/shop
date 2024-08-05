package product

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type DetailParams struct {
	ID uint `json:"id"` // 产品ID
}

// Details 产品详细信息接口
func Details(ctx *context.CustomCtx, params *DetailParams) error {
	data := &detailsInfo{}
	database.Db.Model(&productsModel.Product{}).
		Preload("StoreInfo").Preload("CommentList").Preload("CommentList.UserInfo").
		Preload("FollowInfo", database.Db.Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.StoreFollowStatusConcern)).
		Preload("AttrsList").Preload("AttrsList.Values").Preload("SkuList").
		Where("id = ?", params.ID).
		Find(data)

	return ctx.SuccessJson(data)
}

type detailsInfo struct {
	productsModel.Product
	StoreInfo   shopsModel.Store                   `json:"storeInfo" gorm:"foreignKey:StoreId;"`
	CommentList []storeCommentInfo                 `json:"commentList" gorm:"foreignKey:ProductId;"`
	FollowInfo  shopsModel.StoreFollow             `json:"followInfo" gorm:"foreignKey:ProductId;"`
	AttrsList   []productsModel.ProductAttrsKeyVal `json:"attrsList" gorm:"foreignKey:ProductId;"`
	SkuList     []productsModel.ProductAttrsSku    `json:"skuList" gorm:"foreignKey:ProductId;"`
}

func (_detailsInfo *detailsInfo) TableName() string {
	return "product"
}

type storeCommentInfo struct {
	shopsModel.StoreComment
	UserInfo usersModel.UserInfo `json:"userInfo" gorm:"foreignKey:UserId"`
}

func (storeCommentInfo) TableName() string {
	return "store_comment"
}
