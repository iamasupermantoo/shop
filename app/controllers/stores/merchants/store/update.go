package store

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 店铺更新
type UpdateParams struct {
	Logo     string `json:"logo"`     // Logo
	Name     string `json:"name"`     // 名称
	Address  string `json:"address"`  // 店铺地址
	Contact  string `json:"contact"`  // 联系方式
	Keywords string `json:"keywords"` // 关键词
	Desc     string `json:"desc"`     // 描述
}

// Update 更新店铺
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	database.Db.Model(&shopsModel.Store{}).Where("user_id = ?", ctx.UserId).Updates(&shopsModel.Store{
		Logo:     params.Logo,
		Name:     params.Name,
		Address:  params.Address,
		Contact:  params.Contact,
		Keywords: params.Keywords,
		Desc:     params.Desc,
	})

	return ctx.SuccessJsonOK()
}
