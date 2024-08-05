package setting

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Delete 删除接口
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	for _, id := range params.Ids {
		settingInfo := &adminsModel.AdminSetting{}
		result := database.Db.Model(settingInfo).Where("id = ?", id).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(settingInfo)
		if result.Error == nil && result.RowsAffected > 0 {
			database.Db.Where("id = ?", settingInfo.ID).Delete(&adminsModel.AdminSetting{})
			adminsService.NewAdminSetting(ctx.Rds, settingInfo.AdminId).DelRedisAdminSettingField(settingInfo.Field)
		}
	}

	return ctx.SuccessJsonOK()
}
