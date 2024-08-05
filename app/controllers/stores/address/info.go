package address

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type InfoParams struct {
	ID uint `json:"id" validate:"required"` // 收获地址ID
}

// Info 收获地址详情
func Info(ctx *context.CustomCtx, params *InfoParams) error {
	data := shopsModel.ShippingAddress{}
	database.Db.Model(&shopsModel.ShippingAddress{}).
		Where("id = ?", params.ID).
		Where("user_id = ?", ctx.UserId).
		Where("status = ?", shopsModel.ShippingAddressStatusActivate).
		Find(&data)

	return ctx.SuccessJson(data)
}
