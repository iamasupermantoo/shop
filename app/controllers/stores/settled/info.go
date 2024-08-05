package settled

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Info 商家店铺信息
func Info(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	settledInfo := shopsModel.StoreSettled{}
	result := database.Db.Model(&settledInfo).Where("user_id = ?", ctx.UserId).Find(&settledInfo)
	if result.Error != nil {
		return ctx.ErrorJsonTranslate(result.Error.Error())
	}

	return ctx.SuccessJson(settledInfo)
}
