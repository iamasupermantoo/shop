package category

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

// CreateParams 新增参数
type CreateParams struct {
	ParentId uint   `json:"parentId"`                           //  分类上级ID
	Type     int    `json:"type"`                               //  类型1默认类型
	Name     string `json:"name" validate:"omitempty,alphanum"` //  标题
	Icon     string `json:"icon"`                               //  封面
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	currentAdminId := ctx.AdminId
	if params.ParentId > 0 {
		categoryInfo := &productsModel.Category{}
		result := database.Db.Model(categoryInfo).Where("id = ?", params.ParentId).Find(categoryInfo)
		if result.Error != nil || result.RowsAffected == 0 {
			return ctx.ErrorJson("找不到上级分类信息")
		}
		currentAdminId = categoryInfo.AdminId
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&productsModel.Category{
			ParentId: params.ParentId, AdminId: currentAdminId,
			Type: params.Type, Name: params.Name, Icon: params.Icon,
		})
		if result.Error != nil {
			return ctx.ErrorJson("添加分类失败" + result.Error.Error())
		}

		// 先增对应的翻译数据
		translateInfo := &systemsModel.Translate{}
		result = database.Db.Model(translateInfo).Where("admin_id = ?", currentAdminId).Where("field = ?", params.Name).Find(translateInfo)
		if result.Error != nil {
			return result.Error
		}

		// 如果不存在, 那么新增分类翻译
		if translateInfo.ID == 0 {
			result = tx.Create(&systemsModel.Translate{
				AdminId: currentAdminId,
				Lang:    "zh-CN",
				Field:   params.Name,
				Value:   params.Name,
				Name:    params.Name,
			})
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
