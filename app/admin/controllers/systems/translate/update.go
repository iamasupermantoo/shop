package translate

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID    uint   `gorm:"-" validate:"required" json:"id"` //	ID
	Name  string `json:"name"`                            // 名称
	Lang  string `json:"lang"`                            // 语言标识
	Value string `json:"value"`                           // 键值
	Type  int    `json:"type"`                            // 类型
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	translateInfo := &systemsModel.Translate{}
	result := database.Db.Model(translateInfo).
		Where(params.ID).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Find(translateInfo)
	if result.Error != nil || translateInfo.ID == 0 {
		return ctx.ErrorJson("找不到相应的数据")
	}

	// 更新相应数据并且更新缓存
	database.Db.Model(&systemsModel.Translate{}).
		Where(translateInfo.ID).
		Updates(params)
	systemsService.NewSystemTranslate(ctx.Rds, ctx.AdminSettingId).DelRedisAdminTranslateLangList(translateInfo.Field)

	return ctx.SuccessJsonOK()
}
