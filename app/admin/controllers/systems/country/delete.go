package country

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Delete 删除接口
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	// 删除响应的缓存, 并且删除当前数据
	countryInfo := &systemsModel.Country{}
	if err := database.Db.Model(countryInfo).Where("id IN ?", params.Ids).Where("admin_id IN ?", ctx.GetAdminChildIds()).Delete(countryInfo).Error; err != nil {
		return ctx.ErrorJson(err.Error())
	}
	systemsService.NewSystemCountry(ctx.Rds, ctx.AdminSettingId).DelRedisAdminCountryList()
	return ctx.SuccessJsonOK()
}
