package menu

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	ParentId uint   `json:"parentId"` // 父级ID
	Name     string `json:"name"`     // 名称
	Route    string `json:"route"`    // 路由
	Type     int    `json:"type"`     // 类型1导航菜单 11设置菜单 21更多菜单
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	result := database.Db.Create(&systemsModel.Menu{
		AdminId:  ctx.AdminId,
		ParentId: params.ParentId,
		Name:     params.Name,
		Route:    params.Route,
		Type:     params.Type,
	})
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	adminsService.NewAdminMenu(ctx.Rds, ctx.AdminSettingId).DelRedisAdminMenuList()
	return ctx.SuccessJsonOK()
}
