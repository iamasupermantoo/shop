package comment

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID     uint              `json:"id" gorm:"-" validate:"required"` //	ID
	Rating float64           `json:"rating"`                          //  评分
	Images types.GormStrings `json:"images"`                          //  买家秀
	Name   string            `json:"name"`                            // 评论内容
	Status int               `json:"status"`                          // 状态
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&shopsModel.StoreComment{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
