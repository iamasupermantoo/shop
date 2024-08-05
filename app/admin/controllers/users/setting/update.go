package setting

import (
	"github.com/goccy/go-json"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/usersService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID          uint           `validate:"required"` //	ID
	SettingJson map[string]any `json:"settingJson"`  //	用户配置
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	userInfo := &usersModel.User{}
	database.Db.Model(userInfo).Where("id = ?", params.ID).Find(userInfo)
	userSettingService := usersService.NewUserSetting(ctx.Rds, userInfo.AdminId, userInfo.ID)
	for k, v := range params.SettingJson {
		var value interface{}
		switch v.(type) {
		case map[string]any, []string, []any:
			valueBytes, _ := json.Marshal(v)
			value = string(valueBytes)
		default:
			value = v
		}
		// 更新当前用户设置, 并且删除缓存
		_ = userSettingService.UserSettingUpdate(k, value)
		userSettingService.DelRedisUserSettingField(k)
	}

	return ctx.SuccessJsonOK()
}
