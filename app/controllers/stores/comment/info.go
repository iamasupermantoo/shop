package comment

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type InfoParams struct {
	ID uint `json:"id" validate:"required"` //	产品订单ID
}

type commentProductOrder struct {
	productsModel.ProductOrder
	CommentInfo shopsModel.StoreComment `json:"commentInfo" gorm:"foreignKey:OrderId"`
	ProductInfo productsModel.Product   `json:"productInfo" gorm:"foreignKey:ID;references:ProductId"`
}

func (commentProductOrder) TableName() string {
	return "product_order"
}

// Info 评论信息
func Info(ctx *context.CustomCtx, params *InfoParams) error {
	commentProductOrderInfo := commentProductOrder{}
	database.Db.Model(&productsModel.ProductOrder{}).Where("id = ?", params.ID).Where("user_id = ?", ctx.UserId).
		Preload("CommentInfo").Preload("ProductInfo").Find(&commentProductOrderInfo)

	return ctx.SuccessJson(commentProductOrderInfo)
}
