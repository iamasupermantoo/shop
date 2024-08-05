package shippingCart

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 产品修改购物车参数
type UpdateParams struct {
	ID   uint `json:"id" validate:"required"`        // Id
	Nums int  `json:"nums" validate:"required,gt=0"` // 数量
}

// Update 产品添加到购物车
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	err := database.Db.Model(&shopsModel.StoreCart{}).Where("user_id = ?", ctx.UserId).Where("id = ?", params.ID).Update("nums", params.Nums).Error
	return ctx.IsErrorJson(err)
}
