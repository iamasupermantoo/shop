package auth

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Info 认证信息
func Info(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	data := &usersModel.UserAuth{}
	database.Db.Model(data).Where("user_id = ?", ctx.UserId).Find(data)
	return ctx.SuccessJson(data)
}
