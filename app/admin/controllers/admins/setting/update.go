package setting

import (
	"github.com/goccy/go-json"
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
)

// UpdateParams 更新参数
type UpdateParams struct {
	Items []*adminSetting `json:"items"`
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	for _, item := range params.Items {
		itemValue := item.ValueJson
		switch item.Type {
		case views.InputTypeChildren, views.InputTypeJson, views.InputTypeImages, views.InputTypeCheckbox, views.InputTypeSelect:
			valueBytes, _ := json.Marshal(item.ValueJson)
			itemValue = string(valueBytes)
		}

		settingInfo := &adminsModel.AdminSetting{}
		result := database.Db.Model(settingInfo).Where("id = ?", item.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(settingInfo)
		if result.Error == nil && result.RowsAffected > 0 {
			database.Db.Model(&adminsModel.AdminSetting{}).Where("id = ?", settingInfo.ID).Update("value", itemValue)
			adminsService.NewAdminSetting(ctx.Rds, settingInfo.AdminId).DelRedisAdminSettingField(settingInfo.Field)
		}
	}

	return ctx.SuccessJsonOK()
}
