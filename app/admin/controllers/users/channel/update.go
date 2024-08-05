package channel

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type UpdateParams struct {
	ID     uint   `json:"id" validate:"required"`
	Name   string `json:"name"`                                    //	渠道名称
	Symbol string `json:"symbol"`                                  //	渠道标识
	Route  string `json:"route"`                                   //	请求地址
	Pass   string `json:"pass"`                                    //	密钥
	Mode   int    `json:"mode" validate:"omitempty,oneof=1 11"`    //	模式
	Type   int    `json:"type" validate:"omitempty,oneof=1"`       //	类型
	Status int    `json:"status" validate:"omitempty,oneof=-1 10"` //	状态
}

func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&usersModel.Channel{}).Where("admin_id IN ?", ctx.GetAdminChildIds()).Where("id = ?", params.ID).Updates(params)
	return ctx.IsErrorJson(result.Error)
}
