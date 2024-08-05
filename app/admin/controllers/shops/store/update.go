package store

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID       uint    `json:"id" gorm:"-" validate:"required"` //	ID
	Logo     string  `json:"logo"`                            // logo
	Name     string  `json:"name"`                            // 店铺名称
	Contact  string  `json:"contact"`                         // 联系方式
	Type     int     `json:"type"`                            // 类型
	Keywords string  `json:"keywords"`                        // 关键词
	Desc     string  `json:"desc"`                            // 描述
	Rating   float64 `json:"rating"`                          // 评分
	Score    float64 `json:"score"`                           // 信用分
	Status   int     `json:"status"`                          // 状态 -2删除 -1禁用 10激活 20关店
	Data     string  `json:"data"`                            // 数据
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&shopsModel.Store{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
