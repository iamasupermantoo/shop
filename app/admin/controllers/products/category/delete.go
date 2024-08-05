package category

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

// Delete 删除接口
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		for _, categoryId := range params.Ids {
			categoryInfo := &productsModel.Category{}
			result := database.Db.Model(categoryInfo).Where("id = ?", categoryId).Find(categoryInfo)
			if result.Error == nil && categoryInfo.ID > 0 {
				result = tx.Where("id = ?", categoryInfo.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).Delete(&productsModel.Category{})
				if result.Error != nil {
					return result.Error
				}

				// 删除对应的翻译数据
				result = tx.Where("admin_id = ?", categoryInfo.AdminId).Where("field = ?", categoryInfo.Name).Delete(&systemsModel.Translate{})
				if result.Error != nil {
					return result.Error
				}
			}
		}
		return nil
	})

	return ctx.IsErrorJson(err)
}
