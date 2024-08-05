package role

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

// CreateParams 新增参数
type CreateParams struct {
	Parent string `json:"parent"` // 父级
	Child  string `json:"child"`  // 子级
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		// 列表那边新增一条角色
		authItemInfo := &adminsModel.AuthItem{}
		database.Db.Model(authItemInfo).Where("type = ?", adminsModel.AuthItemTypeRole).Where("name = ?", params.Child).Find(authItemInfo)
		if authItemInfo.ID == 0 {
			authItemInfo = &adminsModel.AuthItem{
				Name: params.Child,
				Type: adminsModel.AuthItemTypeRole,
			}
			result := database.Db.Create(authItemInfo)
			if result.Error != nil {
				return result.Error
			}
		}

		authChildInfo := &adminsModel.AuthChild{}
		database.Db.Model(authChildInfo).Where("parent = ?", params.Parent).Where("child = ?", params.Child).Where("type = ?", adminsModel.AuthChildTypeRoleParentRole).Find(authChildInfo)
		if authChildInfo.ID == 0 {
			authChildInfo = &adminsModel.AuthChild{
				Parent: params.Parent,
				Child:  params.Child,
				Type:   adminsModel.AuthChildTypeRoleParentRole,
			}
			result := tx.Create(authChildInfo)
			if result.Error != nil {
				return result.Error
			}
		}

		return nil
	})
	if err != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + err.Error())
	}

	return ctx.SuccessJsonOK()
}
