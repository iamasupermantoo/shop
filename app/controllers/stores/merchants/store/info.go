package store

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Info 店铺信息
func Info(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	storeInfo := shopsModel.Store{}
	database.Db.Model(&storeInfo).Where("user_id = ?", ctx.UserId).Find(&storeInfo)
	return ctx.SuccessJson(storeInfo)
}
