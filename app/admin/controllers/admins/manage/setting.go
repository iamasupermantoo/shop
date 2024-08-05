package manage

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

type SettingParams struct {
	ID        uint   `validate:"required" json:"id"`        //	ID
	Key       string `validate:"required" json:"key"`       //	授权Key
	Template  string `validate:"required" json:"template"`  //	模版名称
	Nums      int    `validate:"required" json:"nums"`      //	下级数量
	Whitelist string `validate:"required" json:"whitelist"` //	白名单
}

// Setting 管理配置
func Setting(ctx *context.CustomCtx, params *SettingParams) error {
	adminInfo := &adminsModel.AdminUser{}
	result := database.Db.Model(adminInfo).Where("id = ?", params.ID).Where("id IN ?", ctx.GetAdminChildIds()).Find(adminInfo)
	if result.Error != nil || adminInfo.ID == 0 {
		return ctx.ErrorJson("找不到需要修改管理")
	}

	if err := database.Db.Model(adminInfo).Where("id = ?", adminInfo.ID).
		Update("data", utils.JsonMarshal(adminsModel.AdminData{
			Template: params.Template, Key: params.Key,
			AgentNums: params.Nums, Whitelist: params.Whitelist,
		})).Error; err != nil {
		return ctx.ErrorJson(err.Error())
	}

	// 清理管理商户数据
	adminsService.NewAdminUser(ctx.Rds, adminInfo.ID).DelRedisAdminData()
	return ctx.SuccessJsonOK()
}
