package order

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID        uint                 `json:"id" gorm:"-" validate:"required"` // ID
	Money     float64              `json:"money"`                           // 金额
	Status    int                  `json:"status"`                          // 状态-1取消 10等待 11运行 20完成
	ExpiredAt types.GormTimeParams `json:"expiredAt" gorm:"-"`              // 过期时间
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&productsModel.ProductOrder{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)

	return ctx.IsErrorJson(result.Error)
}
