package address

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type DeleteParams struct {
	ID uint `json:"id" validate:"required"` // 收获地址ID
}

// Delete 删除收获地址
func Delete(ctx *context.CustomCtx, params *DeleteParams) error {
	database.Db.Where("id = ?", params.ID).Where("user_id = ?", ctx.UserId).
		Delete(&shopsModel.ShippingAddress{})

	return ctx.SuccessJsonOK()
}
