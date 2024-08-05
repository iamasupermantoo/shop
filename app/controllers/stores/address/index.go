package address

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Index 收获地址列表
func Index(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	data := make([]*shopsModel.ShippingAddress, 0)
	database.Db.Model(&shopsModel.ShippingAddress{}).
		Where("user_id = ?", ctx.UserId).
		Where("status = ?", shopsModel.ShippingAddressStatusActivate).
		Find(&data)

	return ctx.SuccessJson(data)
}
