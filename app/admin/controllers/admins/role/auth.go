package role

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type AuthParams struct {
	Name string `json:"name" validate:"required"` //	权限名称
	Auth string `json:"auth" validate:"required"` //	权限路由
}

func Auth(ctx *context.CustomCtx, params *AuthParams) error {
	authNameInfo := &adminsModel.AuthItem{}
	result := database.Db.Model(authNameInfo).Where("name = ?", params.Name).Where("type = ?", adminsModel.AuthItemTypeName).Find(authNameInfo)
	if result.Error != nil || authNameInfo.ID > 0 {
		return ctx.ErrorJson("权限名称已存在")
	}

	authRouteInfo := &adminsModel.AuthItem{}
	result = database.Db.Model(authRouteInfo).Where("name = ?", params.Auth).Where("type = ?", adminsModel.AuthItemTypeRoute).Find(authRouteInfo)
	if result.Error != nil || authRouteInfo.ID > 0 {
		return ctx.ErrorJson("权限路由已存在")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		result = tx.Create(&adminsModel.AuthItem{Name: params.Name, Type: adminsModel.AuthItemTypeName})
		if result.Error != nil {
			return result.Error
		}
		result = tx.Create(&adminsModel.AuthItem{Name: params.Auth, Type: adminsModel.AuthItemTypeRoute})
		if result.Error != nil {
			return result.Error
		}
		result = tx.Create(&adminsModel.AuthChild{Parent: params.Name, Child: params.Auth, Type: adminsModel.AuthChildTypeRouteNameRoute})
		if result.Error != nil {
			return result.Error
		}

		result = tx.Create(&adminsModel.AuthChild{Parent: adminsModel.AuthRoleSuperManage, Child: params.Name, Type: adminsModel.AuthChildTypeRoleRouteName})
		return result.Error
	})
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}

	// 删除当前管理菜单缓存
	adminsService.NewAdminMenu(ctx.Rds, ctx.AdminId).DelRedisAdminMenuList()
	return ctx.SuccessJsonOK()
}
