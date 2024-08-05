package country

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	Name  string `validate:"required" json:"name"`  // 国家名称
	Alias string `validate:"required" json:"alias"` // 别名
	Iso1  string `validate:"required" json:"iso1"`  // 二位代码
	Icon  string `validate:"required" json:"icon"`  // 图标
	Code  string `validate:"required" json:"code"`  // 区号
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	if err := database.Db.Create(&systemsModel.Country{
		AdminId: ctx.AdminId,
		Name:    params.Name,
		Alias:   params.Alias,
		Icon:    params.Icon,
		Iso1:    params.Iso1,
		Code:    params.Code,
	}).Error; err != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + err.Error())
	}

	// 删除相应的缓存数据
	systemsService.NewSystemCountry(ctx.Rds, ctx.AdminSettingId).DelRedisAdminCountryList()
	return ctx.SuccessJsonOK()
}
