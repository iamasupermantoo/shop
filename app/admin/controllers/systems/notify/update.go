package notify

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID      uint   `gorm:"-" validate:"required" json:"id"` // ID
	Name    string `json:"name"`                            // 标题
	Content string `json:"content"`                         // 内容
	Status  int    `json:"status"`                          // 状态
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	return ctx.IsErrorJson(database.Db.Model(&systemsModel.Notify{}).
		Where("id = ?", params.ID).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params).Error)
}
