package setting

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
	"gofiber/app/module/views"
)

type IndexParams struct {
	GroupId int `json:"groupId"` // 分组ID
}

type adminSetting struct {
	adminsModel.AdminSetting
	DataJson  interface{} `gorm:"-" json:"dataJson"`
	ValueJson interface{} `gorm:"-" json:"valueJson"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*adminSetting, 0)}
	database.Db.Model(&adminsModel.AdminSetting{}).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().Eq("group_id", params.GroupId).Scopes()).
		Find(&data.Items)

	for _, setting := range data.Items.([]*adminSetting) {
		setting.ValueJson = views.InputViewsStringToData(setting.Type, setting.Value)
		setting.DataJson = views.InputViewsStringToData(setting.Type, setting.Data)
	}

	return ctx.SuccessJson(data)
}
