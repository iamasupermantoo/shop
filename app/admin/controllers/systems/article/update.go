package article

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID      uint   `gorm:"-" validate:"required" json:"id"` //	ID
	Image   string `json:"image"`                           // 图片
	Name    string `json:"name"`                            // 标题
	Content string `json:"content"`                         // 内容
	Type    int    `json:"type"`                            // 1基础文章
	Link    string `json:"link"`                            // 链接
	Status  int    `json:"status"`                          // 状态 -1禁用 10开启
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&systemsModel.Article{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)

	return ctx.IsErrorJson(result.Error)
}
