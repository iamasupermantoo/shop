package shippingCart

import (
	shopsModel "gofiber/app/models/service/shopsService"
	"gofiber/app/module/context"
)

// Index 购物车列表
func Index(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	data := shopsModel.NewStoreCar().GetUserStoreCartList(ctx.UserId, []uint{})
	return ctx.SuccessJson(data)
}
