package menu

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	Name   string `json:"name"`   //名称
	Route  string `json:"route"`  //路由
	Status int    `json:"status"` //状态 -1禁用 10开启
	context.IndexParams
}

type adminMenu struct {
	adminsModel.AdminMenu
	ParentInfo *adminsModel.AdminMenu `gorm:"foreignKey:ID;references:ParentId"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*adminMenu, 0)}

	//	过滤参数
	database.Db.Model(&adminsModel.AdminMenu{}).Preload("ParentInfo").
		Scopes(scopes.NewScopes().Eq("name", params.Name).
			Eq("route", params.Route).
			Eq("status", params.Status).
			Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
