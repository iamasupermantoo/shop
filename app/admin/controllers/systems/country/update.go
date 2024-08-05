package country

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 国家更新参数
type UpdateParams struct {
	ID     int    `validate:"required" gorm:"-" json:"id"` // ID
	Name   string `json:"name"`                            // 国家名称
	Alias  string `json:"alias"`                           // 别名
	Icon   string `json:"icon"`                            // 图标
	Iso1   string `json:"iso1"`                            // 二位代码
	Code   string `json:"code"`                            // 区号
	Sort   int    `json:"sort"`                            // 排序
	Status int    `json:"status"`                          // 状态
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	countryInfo := &systemsModel.Country{}
	result := database.Db.Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(countryInfo)
	if result.Error == nil && result.RowsAffected == 0 {
		return ctx.ErrorJson("找不到可更新的数据")
	}

	// 删除相应的缓存
	database.Db.Model(&systemsModel.Country{}).Where("id = ?", countryInfo.ID).Updates(params)
	systemsService.NewSystemCountry(ctx.Rds, ctx.AdminSettingId).DelRedisAdminCountryList()

	return ctx.SuccessJsonOK()
}
