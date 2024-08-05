package channel

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

type CreateParams struct {
	Name   string `json:"name" validate:"required"`            //	渠道名称
	Symbol string `json:"symbol" validate:"required"`          //	渠道标识
	Route  string `json:"route" validate:"required,url"`       //	请求地址
	Mode   int    `json:"mode" validate:"required,oneof=1 11"` //	模式
}

// Create 创建渠道
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	result := database.Db.Create(&usersModel.Channel{
		AdminId: ctx.AdminId, Name: params.Name, Symbol: params.Symbol, Route: params.Route,
		Pass: utils.NewRandom().String(32), Mode: params.Mode,
	})
	return ctx.IsErrorJson(result.Error)
}
