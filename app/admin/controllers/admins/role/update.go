package role

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID     uint   `gorm:"-" validate:"required" json:"id"` //	ID
	Parent string `validate:"required" json:"parent"`      // 父级
	Child  string `validate:"required" json:"child"`       // 子级
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		authChildInfo := &adminsModel.AuthChild{}
		result := database.Db.Model(authChildInfo).Where("id = ?", params.ID).Find(authChildInfo)
		if result.Error != nil || authChildInfo.ID == 0 {
			return ctx.ErrorJson("找不到权限信息")
		}

		// 什么都没有修改, 那么返回
		if params.Child == authChildInfo.Child {
			return nil
		}

		// 先修改auth_item 角色名称
		result = tx.Model(&adminsModel.AuthItem{}).Where("name = ?", authChildInfo.Child).Update("name", params.Child)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Model(&adminsModel.AuthChild{}).
			Where("id = ?", authChildInfo.ID).
			Update("child", params.Child)
		if result.Error != nil {
			return result.Error
		}

		// 清除角色缓存的 - 查询当前角色有哪些管理
		roleAdminList := make([]*adminsModel.AuthAssignment, 0)
		database.Db.Model(&adminsModel.AuthAssignment{}).Where("name = ?", params.Child).Find(&roleAdminList)
		for _, assignment := range roleAdminList {
			adminsService.NewAdminMenu(ctx.Rds, assignment.AdminId).DelRedisAdminMenuList()
		}
		return nil
	})
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}
	return ctx.SuccessJsonOK()
}
