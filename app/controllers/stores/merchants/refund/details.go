package storeRefund

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type DetailsParams struct {
	ID uint `json:"id" validate:"required"` //	产品订单ID
}

type orderRefundInfo struct {
	productsModel.ProductOrder
	ProductInfo productsModel.Product  `json:"productInfo" gorm:"foreignKey:ID;references:ProductId"`
	RefundInfo  shopsModel.StoreRefund `json:"refundInfo" gorm:"foreignKey:ID;references:OrderId"`
	UserInfo    usersModel.UserInfo    `json:"userInfo" gorm:"foreignKey:ID;references:UserId"`
}

func (orderRefundInfo) TableName() string {
	return "product_order"
}

// Details 订单售后详情
func Details(ctx *context.CustomCtx, params *DetailsParams) error {
	storeInfo := &shopsModel.Store{}
	result := database.Db.Model(storeInfo).Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.StoreStatusActivate).Find(storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	refundInfo := orderRefundInfo{}
	database.Db.Model(&productsModel.ProductOrder{}).Preload("ProductInfo").Preload("RefundInfo").Preload("UserInfo").
		Where("id = ?", params.ID).
		Where("store_id = ?", storeInfo.ID).
		Find(&refundInfo)

	return ctx.SuccessJson(refundInfo)
}
