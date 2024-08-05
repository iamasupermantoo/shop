package lang

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID     uint   `gorm:"-" validate:"required" json:"id"` //	ID
	Name   string `json:"name"`                            // 名称
	Alias  string `json:"alias"`                           // 别名
	Icon   string `json:"icon"`                            // 图标
	Sort   int    `json:"sort"`                            // 排序
	Status int    `json:"status"`                          // 状态 -1禁用 10开启
	Symbol string `json:"symbol"`                          // 标识
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	langInfo := &systemsModel.Lang{}
	result := database.Db.Model(langInfo).
		Where("id = ?", params.ID).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Find(langInfo)
	if result.Error != nil {
		return ctx.ErrorJson("没有找到可用数据")
	}

	// 删除相应的缓存数据
	database.Db.Model(&systemsModel.Lang{}).
		Where(langInfo.ID).
		Updates(params)

	systemsService.NewSystemLang(ctx.Rds, ctx.AdminSettingId).DelRedisAdminLangList()
	return ctx.SuccessJsonOK()
}
