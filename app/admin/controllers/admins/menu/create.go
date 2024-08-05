package menu

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	ParentId uint   `json:"parentId"`                 // 父级ID
	Name     string `json:"name" validate:"required"` // 名称
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	if result := database.Db.Create(adminsModel.AdminMenu{
		ParentId: params.ParentId,
		Name:     params.Name,
	}); result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	// 删除所有管理菜单缓存
	adminList := make([]*adminsModel.AdminUser, 0)
	database.Db.Model(&adminsModel.AdminUser{}).Find(&adminList)
	for _, adminInfo := range adminList {
		adminsService.NewAdminMenu(ctx.Rds, adminInfo.ID).DelRedisAdminMenuList()
	}

	return ctx.SuccessJsonOK()
}
