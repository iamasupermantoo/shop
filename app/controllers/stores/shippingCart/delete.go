package shippingCart

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// DeleteParams 产品删除购物车参数
type DeleteParams struct {
	IDs []int `json:"ids" validate:"required"` // 产品Id
}

// Delete 产品添加到购物车
func Delete(ctx *context.CustomCtx, params *DeleteParams) error {
	if result := database.Db.Unscoped().Where("id IN ?", params.IDs).
		Where("user_id = ?", ctx.UserId).Delete(&shopsModel.StoreCart{}); result.Error != nil {
		return ctx.ErrorJsonTranslate(result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
