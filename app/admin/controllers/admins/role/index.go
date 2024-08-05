package role

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
	"gofiber/app/module/views"
)

type IndexParams struct {
	Parent string `json:"parent"` // 父级
	Child  string `json:"child"`  // 子级
	context.IndexParams
}

type authChild struct {
	adminsModel.AuthChild
	AuthList []*views.InputOptions `gorm:"-" json:"authList"` //	权限列表
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*authChild, 0)}

	//	过滤参数
	database.Db.Model(&adminsModel.AuthChild{}).Where("type = ?", adminsModel.AuthChildTypeRoleParentRole).
		Scopes(scopes.NewScopes().
			Eq("parent", params.Parent).
			Eq("child", params.Child).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	authService := adminsService.NewAdminAuth(ctx.Rds, ctx.AdminSettingId)
	for _, childInfo := range data.Items.([]*authChild) {
		childInfo.AuthList = authService.GetAdminRolesRouterSelectOptions([]string{childInfo.Child})
	}

	return ctx.SuccessJson(data)
}
