package invite

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID     uint   `gorm:"-" validate:"required" json:"id"` //	ID
	Code   string `json:"code"`                            // 邀请码
	Status int    `json:"status"`                          // 状态 -1禁用 10开启
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&usersModel.Invite{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
