package menu

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

type SettingParams struct {
	ID      uint   `validate:"required" json:"id"` //	ID
	Icon    string `json:"icon"`                   //	Quasar图标
	ConfURL string `json:"confURL"`                //	配置路由
	Tmp     string `json:"tmp"`                    //	模版名称
}

// Setting 菜单配置
func Setting(ctx *context.CustomCtx, params *SettingParams) error {
	// 更新菜单
	database.Db.Model(&adminsModel.AdminMenu{}).Where("id = ?", params.ID).Update("data", utils.JsonMarshal(adminsModel.AdminMenuData{
		Icon:    params.Icon,
		ConfURL: params.ConfURL,
		Tmp:     params.Tmp,
	}))

	// 删除所有管理菜单
	adminList := make([]*adminsModel.AdminUser, 0)
	database.Db.Model(&adminsModel.AdminUser{}).Find(&adminList)
	for _, adminInfo := range adminList {
		adminsService.NewAdminMenu(ctx.Rds, adminInfo.ID).DelRedisAdminMenuList()
	}
	return ctx.SuccessJsonOK()
}
