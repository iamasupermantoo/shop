package role

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
)

type SettingParams struct {
	ID       uint                  `validate:"required" json:"id"` //	ID
	AuthList []*views.InputOptions `json:"authList"`
}

// Setting 权限配置
func Setting(ctx *context.CustomCtx, params *SettingParams) error {
	authChildInfo := &adminsModel.AuthChild{}
	result := database.Db.Model(authChildInfo).Where("id = ?", params.ID).Where("type = ?", adminsModel.AuthChildTypeRoleParentRole).Find(authChildInfo)
	if result.Error != nil || authChildInfo.ID == 0 {
		return ctx.ErrorJson("找不到权限信息")
	}

	// 删除之前所有角色对应都路由名称
	result = database.Db.Unscoped().Where("parent = ?", authChildInfo.Child).Where("type = ?", adminsModel.AuthChildTypeRoleRouteName).Delete(&adminsModel.AuthChild{})
	if result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	// 添加角色对应都路由名称
	authChildList := make([]*adminsModel.AuthChild, 0)
	for _, auth := range params.AuthList {
		if auth.Value.(bool) {
			authChildList = append(authChildList, &adminsModel.AuthChild{
				Parent: authChildInfo.Child,
				Child:  auth.Label,
				Type:   adminsModel.AuthChildTypeRoleRouteName,
			})
		}
	}
	result = database.Db.Create(&authChildList)
	if result.Error != nil {
		return result.Error
	}

	// 清除角色缓存的 - 查询当前角色有哪些管理
	roleAdminList := make([]*adminsModel.AuthAssignment, 0)
	database.Db.Model(&adminsModel.AuthAssignment{}).Where("name = ?", authChildInfo.Child).Find(&roleAdminList)
	for _, assignment := range roleAdminList {
		adminsService.NewAdminMenu(ctx.Rds, assignment.AdminId).DelRedisAdminMenuList()
	}

	return nil
}
