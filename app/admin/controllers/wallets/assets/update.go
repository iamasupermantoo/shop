package assets

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID     uint    `gorm:"-" validate:"required" json:"id"` //	ID
	Name   string  `json:"name"`                            // 名称
	Symbol string  `json:"symbol"`                          // 标识
	Icon   string  `json:"icon"`                            // 图标
	Type   int     `json:"type"`                            // 类型 1法币资产 11数字货币资产 21虚拟货币资产
	Rate   float64 `json:"rate"`                            // 汇率
	Status int     `json:"status"`                          // 状态 -1禁用 10开启
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&walletsModel.WalletAssets{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
