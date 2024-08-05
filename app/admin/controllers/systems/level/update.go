package level

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID       uint    `gorm:"-" validate:"required" json:"id"` //	ID
	Name     string  `json:"name"`                            // 名称
	Symbol   int     `json:"symbol"`                          // 标识
	Money    float64 `json:"money"`                           //金额
	Icon     string  `json:"icon"`                            // 图标
	Days     int     `json:"days"`                            // 天数
	Type     int     `json:"type"`                            // 购买方式
	Desc     string  `json:"desc"`                            // 详情
	Status   int     `json:"status"`                          // 状态 -1禁用 10开启
	Increase float64 `json:"increase"`                        // 涨幅
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&systemsModel.Level{}).
		Where("id = ?", params.ID).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
