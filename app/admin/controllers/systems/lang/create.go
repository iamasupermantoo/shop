package lang

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	Name   string `validate:"required" json:"name"`   // 名称
	Alias  string `validate:"required" json:"alias"`  // 别名
	Icon   string `validate:"required" json:"icon"`   // 图标
	Symbol string `validate:"required" json:"symbol"` // 标识
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	result := database.Db.Create(&systemsModel.Lang{
		AdminId: ctx.AdminId,
		Name:    params.Name,
		Alias:   params.Alias,
		Icon:    params.Icon,
		Symbol:  params.Symbol,
	})
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	// 删除相应的缓存数据
	systemsService.NewSystemLang(ctx.Rds, ctx.AdminSettingId).DelRedisAdminLangList()

	return ctx.SuccessJsonOK()
}
