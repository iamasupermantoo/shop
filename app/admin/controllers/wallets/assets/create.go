package assets

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	Name   string `validate:"required" json:"name"`   // 名称
	Symbol string `validate:"required" json:"symbol"` // 标识
	Icon   string `validate:"required" json:"icon"`   // 图标
	Type   int    `validate:"required" json:"type"`   // 类型 1法币资产 11数字货币资产 21虚拟货币资产
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	result := database.Db.Create(&walletsModel.WalletAssets{
		AdminId: ctx.AdminId,
		Name:    params.Name,
		Symbol:  params.Symbol,
		Icon:    params.Icon,
		Type:    params.Type,
	})
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
