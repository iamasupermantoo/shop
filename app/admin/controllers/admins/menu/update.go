package menu

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID     uint   `gorm:"-" validate:"required" json:"id"` //	ID
	Name   string `json:"name"`                            // 名称
	Route  string `json:"route"`                           // 路由
	Sort   int    `json:"sort"`                            // 排序
	Status int    `json:"status"`                          // 状态 -1禁用 10开启
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&adminsModel.AdminMenu{}).
		Where("id = ?", params.ID).
		Updates(params)

	// 删除所有管理菜单缓存
	adminList := make([]*adminsModel.AdminUser, 0)
	database.Db.Model(&adminsModel.AdminUser{}).Find(&adminList)
	for _, adminInfo := range adminList {
		adminsService.NewAdminMenu(ctx.Rds, adminInfo.ID).DelRedisAdminMenuList()
	}

	return ctx.IsErrorJson(result.Error)
}
