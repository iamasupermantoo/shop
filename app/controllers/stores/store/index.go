package userStore

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	Search     string             `json:"search"`     // 店铺名称
	Pagination *scopes.Pagination `json:"pagination"` // 分页数据
}

// Index 公用获取店铺列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*shopsModel.Store, 0)}
	database.Db.Model(&shopsModel.Store{}).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			Like("name", "%"+params.Search+"%").
			Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
