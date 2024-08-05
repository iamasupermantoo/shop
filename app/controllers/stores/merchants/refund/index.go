package storeRefund

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

// IndexParams 商家查询售后申请
type IndexParams struct {
	Status     int                `json:"status"`     // 状态
	Pagination *scopes.Pagination `json:"pagination"` // 分页数据
}

// Index 售后列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*storeRefund, 0)}
	storeInfo := &shopsModel.Store{}
	result := database.Db.Model(storeInfo).Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.StoreStatusActivate).Find(storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	database.Db.Model(&shopsModel.StoreRefund{}).Unscoped().Where("store_id = ?", storeInfo.ID).
		Preload("UserInfo").Preload("OrderInfo").Preload("OrderInfo.ProductInfo").
		Scopes(scopes.NewScopes().
			Eq("status", params.Status).Scopes()).Scopes(params.Pagination.Scopes()).
		Count(&data.Count).Find(&data.Items)

	return ctx.SuccessJson(data)
}

type storeRefund struct {
	shopsModel.StoreRefund
	UserInfo  usersModel.UserInfo `json:"userInfo" gorm:"foreignKey:ID;references:UserId"`
	OrderInfo productOrder        `json:"orderInfo" gorm:"foreignKey:ID;references:OrderId"`
}

func (storeRefund) TableName() string {
	return "store_refund"
}

type productOrder struct {
	productsModel.ProductOrder
	ProductInfo productsModel.Product `json:"productInfo" gorm:"foreignKey:ID;references:ProductId"`
}
